package globals

var ReflectionLookup = map[int]ReflectionItem{}

type ReflectionItem struct {
    Title string
    Function interface{}
}
