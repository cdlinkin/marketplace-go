package repo

func NewProductRepo(kind string) ProductRepo {
	switch kind {
	case "memory":
		return NewMemoryProductRepo()
	case "file":
		return NewFileProductRepo("products.json")
	}
	return nil
}

func NewOrderRepo(kind string) OrderRepo {
	switch kind {
	case "memory":
		return NewMemoryOrderRepo()
	case "file":
		return NewFileOrderRepo("orders.json")
	}
	return nil
}

func NewUserRepo(kind string) UserRepo {
	switch kind {
	case "memory":
		return NewMemoryUserRepo()
	case "file":
		return NewFileUserRepo("users.json")
	}
	return nil
}
