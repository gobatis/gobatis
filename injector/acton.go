package injector

type Injector interface {
}

func SoftDelete(field string) Injector {
	return nil
}

func OptimisticLock(field string) Injector {
	return nil
}

func SetUpdatedAt(field string) Injector {
	return nil
}
