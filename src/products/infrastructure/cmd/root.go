package products_cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"

	products_application "github.com/juanfer2/shopy_dc_golang/src/products/application"
	products_domain "github.com/juanfer2/shopy_dc_golang/src/products/domain"
	products_repositories "github.com/juanfer2/shopy_dc_golang/src/products/infrastructure/repositories"
)

var PersonaSeedCmd = &cobra.Command{
	Use:     "db:seed",
	Aliases: []string{"m"},
	Short:   "Run Seeds for products",
}

func init() {
	PersonaSeedCmd.AddCommand(runSeedCmd)
}

var runSeedCmd = &cobra.Command{
	Use:     "products",
	Aliases: []string{"m"},
	Short:   "Run seeds.",
	Run:     runSeed,
}

func runSeed(cmd *cobra.Command, args []string) {
	color.Blue(">>>Create Data...  🚀 <<<")
	products := []products_domain.Product{
		{Name: "Apple Air Pods 2", Price: 600, ImageUrl: "airpods.png"},
		{Name: "Samsung Smart Watch", Price: 600, ImageUrl: "smart_watch.png"},
		{Name: "Apple Monitor Pro", Price: 600, ImageUrl: "tv.png"},
		{Name: "Google Pixel 4 XL", Price: 600, ImageUrl: "samsung_phone.png"},
		{Name: "Amazon Smart Speaker", Price: 600, ImageUrl: "smart_speaker.png"},
		{Name: "Apple MacBook Pro 16", Price: 600, ImageUrl: "macbook_air.png"},
		{Name: "Microsoft Surface Book", Price: 600, ImageUrl: "surface_book.png"},
		{Name: "Google Nest", Price: 600, ImageUrl: "google_nest.png"},
		{Name: "Apple iMac Pro", Price: 600, ImageUrl: "apple_ipad.png"},
	}

	productRepository := products_repositories.NewProductPostgresRepository()
	productService := products_application.NewProductUseCase(productRepository)
	productService.CreateMultiplesProducts(products)
	color.Blue(">>>Success migration... 🚀 <<<")
}
