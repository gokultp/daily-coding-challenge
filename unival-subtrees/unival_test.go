package unival

import "testing"

func TestTree_CountUnivalSubtrees(t *testing.T) {
	type fields struct {
		Root *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{name: "test", fields: fields {
			Root: &Node{
				Value : 0,
				Left : &Node{Value :1},
				Right : &Node{
					Value : 0,
					Left :&Node {
						Value : 1,
						Left : &Node{Value : 1},
						Right : &Node{Value : 1},
					},
					Right :&Node {Value : 0},
				},
			},
		}, want : 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tree{
				Root: tt.fields.Root,
			}
			if got := tr.CountUnivalSubtrees(); got != tt.want {
				t.Errorf("Tree.CountUnivalSubtrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

