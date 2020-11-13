package main

import "fmt"

type Book struct{
	ID int
	Name string
	Author string
	YearPublished int
}

func(b Book) String() string {
	return fmt.Sprintf("\n\tBook :%v \n"+
		"\tID:%v \n"+
		"\tAuthor:%v",b.Name,b.ID,b.Author)
}


var books =[]Book{
		Book{
		ID:            1,
		Name:          "The First",
		Author:        "Author1",
		YearPublished: 1990,
		},
		Book{
		ID:            2,
		Name:         "The second",
		Author:        "Mr second",
		YearPublished: 1998,
		},
		Book{
		ID:            3,
		Name:         "The Third",
		Author:        "Mr. Third",
		YearPublished: 1999,
		},
		Book{
		ID:            4,
		Name:         "The Fourth",
		Author:        "Mr. Fourth",
		YearPublished: 1991,
		},
		Book{
		ID:            5,
		Name:         "The Five",
		Author:        "Mr. Five",
		YearPublished: 2001,
		},
		Book{
		ID:            6,
		Name:         "The Six",
		Author:        "Mr. Six",
		YearPublished: 2002,
		},
		Book{
		ID:            7,
		Name:         "The Seven",
		Author:        "Mr. Seven",
		YearPublished: 2003,
		},
}