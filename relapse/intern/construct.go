//  Copyright 2017 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package intern

import (
	"fmt"

	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/compose"
	"github.com/katydid/katydid/relapse/funcs"
	"github.com/katydid/katydid/relapse/interp"
	nameexpr "github.com/katydid/katydid/relapse/name"
)

type Construct interface {
	AddGrammar(g *ast.Grammar) (*Pattern, error)
	AddPatternDecl(name string, p *ast.Pattern) (*Pattern, error)
	Deref(name string) *Pattern
	NewPattern(this *ast.Pattern) (*Pattern, error)
	NewConcat(ps []*Pattern) *Pattern
	NewOr(ps []*Pattern) (*Pattern, error)
	NewAnd(ps []*Pattern) (*Pattern, error)
	NewInterleave(ps []*Pattern) *Pattern
	NewZAny() *Pattern
	NewEmpty() *Pattern
	NewNot(*Pattern) *Pattern
	NewZeroOrMore(*Pattern) *Pattern
}

type construct struct {
	refs     map[string]*Pattern
	hashRefs map[uint64][]string
	nullRefs map[string]bool
	cache    map[uint64][]*Pattern
	record   bool
}

func NewConstructor() Construct {
	return &construct{
		refs:     make(map[string]*Pattern),
		hashRefs: make(map[uint64][]string),
		nullRefs: make(map[string]bool),
		cache:    make(map[uint64][]*Pattern),
	}
}

func NewConstructorOptimizedForRecords() Construct {
	return &construct{
		refs:     make(map[string]*Pattern),
		hashRefs: make(map[uint64][]string),
		nullRefs: make(map[string]bool),
		cache:    make(map[uint64][]*Pattern),
		record:   true,
	}
}

func (c *construct) AddPatternDecl(name string, p *ast.Pattern) (*Pattern, error) {
	pp, err := c.NewPattern(p)
	if err != nil {
		return nil, err
	}
	c.refs[name] = pp
	c.hashRefs[pp.hash] = append(c.hashRefs[pp.hash], name)
	return pp, nil
}

func (c *construct) Deref(name string) *Pattern {
	return c.refs[name]
}

func (c *construct) AddGrammar(g *ast.Grammar) (*Pattern, error) {
	ps := g.GetPatternDecls()
	refs := ast.NewRefLookup(g)
	c.nullRefs = make(map[string]bool)
	for name, p := range refs {
		c.nullRefs[name] = interp.Nullable(refs, p)
	}

	main, err := c.AddPatternDecl("main", g.GetTopPattern())
	if err != nil {
		return nil, err
	}

	for _, p := range ps {
		if _, err := c.AddPatternDecl(p.GetName(), p.GetPattern()); err != nil {
			return nil, err
		}
	}
	return main, nil
}

func (c *construct) checkRef(p *Pattern) (*Pattern, error) {
	nn := c.hashRefs[p.hash]
	for _, name := range nn {
		if c.refs[name].Equal(p) {
			return c.NewReference(name)
		}
	}
	return p, nil
}

func (c *construct) NewPattern(this *ast.Pattern) (*Pattern, error) {
	if this.Empty != nil {
		return c.NewEmpty(), nil
	}
	if this.TreeNode != nil {
		b := nameexpr.NameToFunc(this.TreeNode.GetName())
		b = funcs.Simplify(b)
		p, err := c.NewPattern(this.TreeNode.GetPattern())
		if err != nil {
			return nil, err
		}
		return c.NewNode(b, p)
	}
	if this.LeafNode != nil {
		b, err := compose.NewBool(this.LeafNode.GetExpr())
		if err != nil {
			return nil, err
		}
		b = funcs.Simplify(b)
		return c.NewNode(b, empty)
	}
	if this.Concat != nil {
		concats := getConcats(this)
		ps, err := deriveTraverse(c.NewPattern, concats)
		if err != nil {
			return nil, err
		}
		return c.NewConcat(ps), nil
	}
	if this.Or != nil {
		ors := getOrs(this)
		ps, err := deriveTraverse(c.NewPattern, ors)
		if err != nil {
			return nil, err
		}
		return c.NewOr(ps)
	}
	if this.And != nil {
		ands := getAnds(this)
		ps, err := deriveTraverse(c.NewPattern, ands)
		if err != nil {
			return nil, err
		}
		return c.NewAnd(ps)
	}
	if this.ZeroOrMore != nil {
		p, err := c.NewPattern(this.ZeroOrMore.GetPattern())
		if err != nil {
			return nil, err
		}
		return c.NewZeroOrMore(p), nil
	}
	if this.Reference != nil {
		return c.NewReference(this.Reference.GetName())
	}
	if this.Not != nil {
		p, err := c.NewPattern(this.Not.GetPattern())
		if err != nil {
			return nil, err
		}
		return c.NewNot(p), nil
	}
	if this.ZAny != nil {
		return c.NewZAny(), nil
	}
	if this.Contains != nil {
		p, err := c.NewPattern(this.Contains.GetPattern())
		if err != nil {
			return nil, err
		}
		return c.NewContains(p), nil
	}
	if this.Optional != nil {
		p, err := c.NewPattern(this.Optional.GetPattern())
		if err != nil {
			return nil, err
		}
		return c.NewOptional(p), nil
	}
	if this.Interleave != nil {
		interleaves := getInterleaves(this)
		ps, err := deriveTraverse(c.NewPattern, interleaves)
		if err != nil {
			return nil, err
		}
		return c.NewInterleave(ps), nil
	}
	return nil, fmt.Errorf("unknown pattern: %v", this)
}

func getConcats(p *ast.Pattern) []*ast.Pattern {
	if p.Concat != nil {
		return append(getConcats(p.Concat.GetLeftPattern()), getConcats(p.Concat.GetRightPattern())...)
	}
	return []*ast.Pattern{p}
}

func getOrs(p *ast.Pattern) []*ast.Pattern {
	if p.Or != nil {
		return append(getOrs(p.Or.GetLeftPattern()), getOrs(p.Or.GetRightPattern())...)
	}
	return []*ast.Pattern{p}
}

func getAnds(p *ast.Pattern) []*ast.Pattern {
	if p.And != nil {
		return append(getAnds(p.And.GetLeftPattern()), getAnds(p.And.GetRightPattern())...)
	}
	return []*ast.Pattern{p}
}

func getInterleaves(p *ast.Pattern) []*ast.Pattern {
	if p.Interleave != nil {
		return append(getInterleaves(p.Interleave.GetLeftPattern()), getInterleaves(p.Interleave.GetRightPattern())...)
	}
	return []*ast.Pattern{p}
}

var empty = &Pattern{
	Type:     Empty,
	nullable: true,
}

var zany = &Pattern{
	Type:     ZAny,
	nullable: true,
}

var notzany = &Pattern{
	Type:     Not,
	Patterns: []*Pattern{zany},
	nullable: false,
}

func init() {
	empty.SetHash()
	zany.SetHash()
	notzany.SetHash()
}

func (c *construct) NewEmpty() *Pattern {
	return empty
}

func (c *construct) NewZAny() *Pattern {
	return zany
}

func (c *construct) NewNotZAny() *Pattern {
	return notzany
}

func (c *construct) NewNode(b funcs.Bool, child *Pattern) (*Pattern, error) {
	if child == notzany {
		return notzany, nil
	}
	if funcs.IsFalse(b) {
		return notzany, nil
	}
	p := &Pattern{
		Type:     Node,
		Func:     b,
		Patterns: []*Pattern{child},
		nullable: false,
	}
	p.SetHash()
	return c.checkRef(p)
}

func isNotZAny(p *Pattern) bool {
	return p == notzany
}

func isZany(p *Pattern) bool {
	return p == zany
}

func notEmptySet(p *Pattern) bool { return p != notzany }

func removeNotZAny(ps []*Pattern) []*Pattern {
	return deriveFilter(notEmptySet, ps)
}

func notEmpty(p *Pattern) bool { return p != empty }

func removeEmpty(ps []*Pattern) []*Pattern {
	return deriveFilter(notEmpty, ps)
}

func isNullable(p *Pattern) bool {
	return p.nullable
}

func (c *construct) MergeOr(l, r *Pattern) (*Pattern, error) {
	var left []*Pattern
	if l.Type == Or {
		left = l.Patterns
	} else {
		left = []*Pattern{l}
	}
	var right []*Pattern
	if r.Type == Or {
		right = r.Patterns
	} else {
		right = []*Pattern{r}
	}
	return c.NewOr(append(left, right...))
}

func (c *construct) NewOr(ps []*Pattern) (*Pattern, error) {
	if deriveAny(isZany, ps) {
		return zany, nil
	}
	ps = removeNotZAny(ps)
	if deriveAll(isNullable, ps) {
		ps = removeEmpty(ps)
	}
	if len(ps) == 0 {
		return zany, nil
	}
	ps = orderedSet(ps)
	var err error
	ps, err = c.mergeLeaves(funcs.Or, ps)
	if err != nil {
		return nil, err
	}
	if c.record {
		ps, err = c.mergeContainsOr(ps)
		if err != nil {
			return nil, err
		}
	}
	ps, err = c.mergeNodesOr(ps)
	if err != nil {
		return nil, err
	}
	pp := &Pattern{
		Type:     Or,
		Patterns: ps,
		nullable: anyNullable(ps),
	}
	pp.SetHash()
	return pp, nil
}

func merge(predicate func(l, r *Pattern) bool, merge func(l, r *Pattern) (*Pattern, error), ps []*Pattern) ([]*Pattern, error) {
	if len(ps) == 0 {
		return nil, nil
	}
	last := 0
	for i := 1; i < len(ps); i++ {
		if !predicate(ps[last], ps[i]) {
			last++
			continue
		}
		var err error
		ps[last], err = merge(ps[last], ps[i])
		if err != nil {
			return nil, err
		}
	}
	return ps[:last+1], nil
}

func areLeaves(l, r *Pattern) bool {
	return l.Type == Node && l.Patterns[0] == empty &&
		r.Type == Node && r.Patterns[0] == empty
}

func (c *construct) mergeLeaves(mergeFunc func(l, r funcs.Bool) funcs.Bool, ps []*Pattern) ([]*Pattern, error) {
	return merge(areLeaves, func(l, r *Pattern) (*Pattern, error) {
		return c.NewNode(mergeFunc(l.Func, r.Func), c.NewEmpty())
	}, ps)
}

func areContainsWithEqualNames(l, r *Pattern) bool {
	return l.Type == Contains && r.Type == Contains &&
		l.Patterns[0].Type == Node && r.Patterns[0].Type == Node &&
		funcs.Equal(l.Patterns[0].Func, r.Patterns[0].Func)
}

func (c *construct) mergeContainsOr(ps []*Pattern) ([]*Pattern, error) {
	return merge(areContainsWithEqualNames, func(l, r *Pattern) (*Pattern, error) {
		o, err := c.MergeOr(l.Patterns[0], r.Patterns[0])
		if err != nil {
			return nil, err
		}
		pp, err := c.NewNode(l.Func, o)
		if err != nil {
			return nil, err
		}
		return c.NewContains(pp), nil
	}, ps)
}

func areNodesWithEqualNames(l, r *Pattern) bool {
	return l.Type == Node && r.Type == Node && funcs.Equal(l.Func, r.Func)
}

func (c *construct) mergeNodesOr(ps []*Pattern) ([]*Pattern, error) {
	return merge(areNodesWithEqualNames, func(l, r *Pattern) (*Pattern, error) {
		o, err := c.MergeOr(l.Patterns[0], r.Patterns[0])
		if err != nil {
			return nil, err
		}
		return c.NewNode(l.Func, o)
	}, ps)
}

func (c *construct) NewConcat(ps []*Pattern) *Pattern {
	if deriveAny(isNotZAny, ps) {
		return notzany
	}
	ps = removeEmpty(ps)
	if len(ps) == 3 {
		if ps[0] == zany && ps[2] == zany {
			return c.NewContains(ps[1])
		}
	}
	pp := &Pattern{
		Type:     Concat,
		Patterns: ps,
		nullable: allNullable(ps),
	}
	pp.SetHash()
	return pp
}

func removeZAny(ps []*Pattern) []*Pattern {
	return deriveFilter(func(p *Pattern) bool {
		return p.Type != ZAny
	}, ps)
}

func isEmpty(p *Pattern) bool {
	return p == empty
}

func isNotZany(p *Pattern) bool {
	return p == notzany
}

func (c *construct) MergeAnd(l, r *Pattern) (*Pattern, error) {
	var left []*Pattern
	if l.Type == Or {
		left = l.Patterns
	} else {
		left = []*Pattern{l}
	}
	var right []*Pattern
	if r.Type == Or {
		right = r.Patterns
	} else {
		right = []*Pattern{r}
	}
	return c.NewAnd(append(left, right...))
}

func (c *construct) NewAnd(ps []*Pattern) (*Pattern, error) {
	if deriveAny(isNotZany, ps) {
		return notzany, nil
	}
	ps = removeNotZAny(ps)
	if deriveAny(isEmpty, ps) {
		if deriveAll(isNullable, ps) {
			return empty, nil
		}
		return notzany, nil
	}
	if len(ps) == 0 {
		return zany, nil
	}
	ps = orderedSet(ps)
	var err error
	ps, err = c.mergeLeaves(funcs.And, ps)
	if err != nil {
		return nil, err
	}
	if c.record {
		ps, err = c.mergeContainsAnd(ps)
		if err != nil {
			return nil, err
		}
	}
	ps, err = c.mergeNodesAnd(ps)
	if err != nil {
		return nil, err
	}
	pp := &Pattern{
		Type:     And,
		Patterns: ps,
		nullable: allNullable(ps),
	}
	pp.SetHash()
	return pp, nil
}

func (c *construct) mergeContainsAnd(ps []*Pattern) ([]*Pattern, error) {
	return merge(areContainsWithEqualNames, func(l, r *Pattern) (*Pattern, error) {
		a, err := c.MergeAnd(l.Patterns[0], r.Patterns[0])
		if err != nil {
			return nil, err
		}
		pp, err := c.NewNode(l.Func, a)
		if err != nil {
			return nil, err
		}
		return c.NewContains(pp), nil
	}, ps)
}

func (c *construct) mergeNodesAnd(ps []*Pattern) ([]*Pattern, error) {
	return merge(areNodesWithEqualNames, func(l, r *Pattern) (*Pattern, error) {
		a, err := c.MergeAnd(l.Patterns[0], r.Patterns[0])
		if err != nil {
			return nil, err
		}
		return c.NewNode(l.Func, a)
	}, ps)
}

func (c *construct) NewZeroOrMore(p *Pattern) *Pattern {
	if p.Type == ZeroOrMore {
		return p
	}
	pp := &Pattern{
		Type:     ZeroOrMore,
		Patterns: []*Pattern{p},
		nullable: true,
	}
	pp.SetHash()
	return pp
}

func (c *construct) NewReference(name string) (*Pattern, error) {
	n, ok := c.nullRefs[name]
	if !ok {
		return nil, fmt.Errorf("no reference with name: %s", name)
	}
	pp := &Pattern{
		Type:     Reference,
		Ref:      name,
		nullable: n,
	}
	pp.SetHash()
	return pp, nil
}

func (c *construct) NewNot(p *Pattern) *Pattern {
	if p == zany {
		return notzany
	}
	if p.Type == Not {
		return p.Patterns[0]
	}
	pp := &Pattern{
		Type:     Not,
		Patterns: []*Pattern{p},
		nullable: !p.nullable,
	}
	pp.SetHash()
	return pp
}

func (c *construct) NewContains(p *Pattern) *Pattern {
	if p == empty || p == zany {
		return zany
	}
	if p == notzany {
		return notzany
	}
	pp := &Pattern{
		Type:     Contains,
		Patterns: []*Pattern{p},
		nullable: p.nullable,
	}
	pp.SetHash()
	return pp
}

func (c *construct) NewOptional(p *Pattern) *Pattern {
	if p == empty {
		return empty
	}
	pp := &Pattern{
		Type:     Optional,
		Patterns: []*Pattern{p},
		nullable: true,
	}
	pp.SetHash()
	return pp
}

func (c *construct) NewInterleave(ps []*Pattern) *Pattern {
	if deriveAny(isNotZAny, ps) {
		return notzany
	}
	ps = removeEmpty(ps)
	ps = removeZAny(ps)
	ps = orderedSet(ps)
	pp := &Pattern{
		Type:     Interleave,
		Patterns: ps,
		nullable: allNullable(ps),
	}
	pp.SetHash()
	return pp
}
