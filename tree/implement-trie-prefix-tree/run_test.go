package implement_trie_prefix_tree

import (
	"reflect"
	"testing"
)

func TestTrie_Search(t1 *testing.T) {
	type fields struct {
		children [26]*Trie
		isEnd    bool
	}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "test1",
			fields: fields{},
			args: args{
				word: "abccde",
			},
			want: true,
		},
	}
	node := Constructor()
	node.Insert("abc")
	tests[0].fields = fields{children: node.children, isEnd: node.isEnd}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Trie{
				children: tt.fields.children,
				isEnd:    tt.fields.isEnd,
			}
			if got := t.Search(tt.args.word); got != tt.want {
				t1.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrie_SearchPrefix(t1 *testing.T) {
	type fields struct {
		children [26]*Trie
		isEnd    bool
	}
	type args struct {
		prefix string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Trie
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Trie{
				children: tt.fields.children,
				isEnd:    tt.fields.isEnd,
			}
			if got := t.SearchPrefix(tt.args.prefix); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("SearchPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrie_StartsWith(t1 *testing.T) {
	type fields struct {
		children [26]*Trie
		isEnd    bool
	}
	type args struct {
		prefix string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{{
		name:   "test1",
		fields: fields{},
		args: args{
			prefix: "abc",
		},
		want: true,
	},
	}
	node := Constructor()
	node.Insert("abccde")
	tests[0].fields = fields{children: node.children, isEnd: node.isEnd}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Trie{
				children: tt.fields.children,
				isEnd:    tt.fields.isEnd,
			}
			if got := t.StartsWith(tt.args.prefix); got != tt.want {
				t1.Errorf("StartsWith() = %v, want %v", got, tt.want)
			}
		})
	}
}
