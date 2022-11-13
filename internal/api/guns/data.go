package guns

import Shared "go-api/internal/shared"

var Guns = []Shared.Model[IGun]{
	{Data: IGun{Id: "1", Name: "AK-47", Price: 2000}, Hash: "guns:1"},
	{Data: IGun{Id: "2", Name: "Glock", Price: 500}, Hash: "guns:2"},
	{Data: IGun{Id: "3", Name: "MP5", Price: 1150}, Hash: "guns:3"},
}
