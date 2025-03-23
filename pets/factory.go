package pets

import "patterns/models"

func New(species string) *models.Pet {
	pet := models.Pet{
		Species:     species,
		Breed:       "",
		MinWeight:   0,
		MaxWeight:   0,
		Description: "no description entered yet",
		LifeSpan:    0,
	}

	return &pet
}
