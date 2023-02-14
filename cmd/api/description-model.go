package main

type DescriptionModel struct {
	descriptions []string
}

func createDescriptionModel() DescriptionModel {
	var descriptions []string

	descriptions = []string{"Dank Pirates", "Gay Harambe NFT's", "Description of Dank Project", "Is Badass"}

	return DescriptionModel{
		descriptions: descriptions,
	}
}
