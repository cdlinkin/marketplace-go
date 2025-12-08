package repo

func NewProductRepo(kind string) ProductRepo {
	switch kind {
	case "memory":
		return NewMemoryProductRepo()
	}
	return nil
}

func NewOrderRepo(kind string) OrderRepo {
	switch kind {
	case "memory":
		return NewMemoryOrderRepo()
	}
	return nil
}

func NewUserRepo(kind string) UserRepo {
	switch kind {
	case "memory":
		return NewMemoryUserRepo()
	}
	return nil
}
