package docs

func SetSwagger() {
	SwaggerInfo.Title = "Warehouse API"
	SwaggerInfo.Description = "This is warehouse server"
	SwaggerInfo.Version = "1.0"
	SwaggerInfo.Host = "localhost:8088"
	SwaggerInfo.BasePath = "/"
	SwaggerInfo.Schemes = []string{"http"}
}
