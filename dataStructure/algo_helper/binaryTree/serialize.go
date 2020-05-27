package binaryTree

type Serializer interface {
	Serialize() string
}

type Deserializer interface {
	Deserialize(string) *Node
}

type SerializeDeserializer interface {
	Serializer
	Deserializer
}

type LevelOrderSerializeOrDe struct {
	root *Node
}

func (l *LevelOrderSerializeOrDe) Serialize() string {
	return ""
}

func (l *LevelOrderSerializeOrDe) Deserialize(input string) *Node {
	return nil
}