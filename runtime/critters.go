package runtime

type Critters interface {
	Update(AppContext)
	Draw(AppContext)
}
