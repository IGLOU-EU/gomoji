// Package gomoji implements the list of emoji officially available from the
// Unicode Consortium.
//
// You can access directly to the array to perform your operations or use the
// utility functions implemented with this library.
package gomoji

type Emojis []Emoji

// An Emoji serves the necessary information for the use of emojis as its category
// and its keywords.
type Emoji struct {
	Picto    string
	Name     string
	Category string
	Keywords []string
}

type Groups []Group

type Group struct {
	Name string
	Size int
}

// CategorySize return the number of emoji on a category
// or -1 if the category does not exist
func (g *Groups) Size(name string) int {
	for i := range *g {
		if (*g)[i].Name == name {
			return (*g)[i].Size
		}
	}

	return -1
}

// Get try to find an emoji by its name.
// Returns a pointer to an Emoji or nill
func Get(name string) *Emoji {
	for i := range List {
		if List[i].Name == name {
			return &List[i]
		}
	}

	return nil
}

// Info return the emoji information if exist
// Returns a pointer to an Emoji or nill
func Info(picto string) *Emoji {
	for i := range List {
		if List[i].Picto == picto {
			return &List[i]
		}
	}

	return nil
}

// ByKeyword return a slice of Emoji coresponding to keyword if it exist
// Returns a pointer to an Emoji or nill
func ByKeyword(keyword string) Emojis {
	size := ListKeyword.Size(keyword)
	if size <= 0 {
		return Emojis{}
	}

	emojis := make(Emojis, 0, size)
	for i := range List {
		for _, k := range List[i].Keywords {
			if k == keyword {
				emojis = append(emojis, List[i])
			}
		}
	}

	return emojis

}

// ByCategory return a slice of Emoji from this category if it exist
// or an empty slice
func ByCategory(category string) Emojis {
	size := ListCategory.Size(category)
	if size <= 0 {
		return Emojis{}
	}

	emojis := make(Emojis, 0, size)
	for i := range List {
		if List[i].Category == category {
			emojis = append(emojis, List[i])
		}
	}

	return emojis
}

// Categorysed return a map where id are the category name and value a slice
// of there emojis
func Categorysed() map[string]Emojis {
	categorysed := make(map[string]Emojis, ListCategorySize)

	for i := range ListCategory {
		name := ListCategory[i].Name
		categorysed[name] = ByCategory(name)
	}

	return categorysed
}
