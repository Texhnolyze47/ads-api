package router

import (
	"comercial-api/controller"
	"comercial-api/db"
	"comercial-api/internal/database"
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Pedido struct {
	Cantidad    float64   `json:"cantidad"`
	Fecha       time.Time `json:"fecha"`
	IDCliente   uuid.UUID `json:"id_cliente"`
	IDComercial uuid.UUID `json:"id_comercial"`
}

type Comercial struct {
	Nombre          string  `json:"nombre"`
	ApellidoPaterno string  `json:"apellido_paterno"`
	ApellidoMaterno string  `json:"apellido_materno"`
	Ciudad          string  `json:"ciudad"`
	Commision       float64 `json:"commision"`
}

func SetupRouter(cfg db.ApiConfig, e *echo.Echo) (*echo.Echo, error) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/cliente", controller.HandlerCreateClient(cfg))
	e.POST("/pedido", func(c echo.Context) error {

		var newOrder Pedido

		err := c.Bind(&newOrder)

		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		order, err := cfg.DB.CreateOrder(context.Background(), database.CreateOrderParams{
			ID:          uuid.New(),
			Cantidad:    newOrder.Cantidad,
			Fecha:       newOrder.Fecha,
			IDCliente:   newOrder.IDCliente,
			IDComercial: newOrder.IDComercial,
		})

		if err != nil {
			return c.String(500, err.Error())
		}

		return c.JSON(200, order)
	})

	e.POST("/comercial", func(c echo.Context) error {
		var newComercial Comercial

		err := c.Bind(&newComercial)
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad request")
		}

		comercial, err := cfg.DB.CreateCommercial(context.Background(), database.CreateCommercialParams{
			ID:              uuid.New(),
			Nombre:          newComercial.Nombre,
			ApellidoPaterno: newComercial.ApellidoPaterno,
			ApellidoMaterno: newComercial.ApellidoMaterno,
			Ciudad:          newComercial.Ciudad,
			Commision:       newComercial.Commision,
		})

		if err != nil {
			return c.String(500, err.Error())

		}

		return c.JSON(200, comercial)

	})

	e.GET("/consultas/1", listarPedidos(cfg))
	e.GET("/consultas/2", listarDosPedidosMayorValor(cfg))
	e.GET("/consultas/3", listarIDsClientes(cfg))
	e.GET("/consultas/4", listarPedidosAnio2017(cfg))
	e.GET("/consultas/5", listarComercialesComision(cfg))
	e.GET("/consultas/6", obtenerMayorComision(cfg))
	e.GET("/consultas/7", listarClientesSegundoApellido(cfg))
	e.GET("/consultas/8", listarNombresEspecificos(cfg))
	e.GET("/consultas/9", listarNombresSinA(cfg))
	e.GET("/consultas/10", listarNombresComerciales(cfg))

	e.Logger.Fatal(e.Start(":1323"))

	return e, nil
}

func listarNombresComerciales(cfg db.ApiConfig) echo.HandlerFunc {
	return func(c echo.Context) error {

		listNameAd, err := cfg.DB.ONameADName(context.Background())

		if err != nil {
			c.String(500, err.Error())
		}

		return c.JSON(200, listNameAd)

	}
}

func listarNombresSinA(cfg db.ApiConfig) echo.HandlerFunc {
	return func(c echo.Context) error {
		listNameNoA, err := cfg.DB.ANoStartName(context.Background())

		if err != nil {
			return c.String(500, err.Error())
		}

		return c.JSON(200, listNameNoA)
	}

}

func listarNombresEspecificos(cfg db.ApiConfig) echo.HandlerFunc {
	return func(c echo.Context) error {

		listNameStart, err := cfg.DB.AStartName(context.Background())

		if err != nil {
			return c.String(500, err.Error())
		}

		return c.JSON(200, listNameStart)
	}

}

func listarClientesSegundoApellido(cfg db.ApiConfig) echo.HandlerFunc {
	return func(c echo.Context) error {
		listLastName, err := cfg.DB.LastNameNotNull(context.Background())

		if err != nil {
			return c.String(500, err.Error())
		}

		return c.JSON(200, listLastName)

	}

}

func obtenerMayorComision(cfg db.ApiConfig) echo.HandlerFunc {
	return func(c echo.Context) error {

		listMaxCommision, err := cfg.DB.MaxCommission(context.Background())

		if err != nil {
			return c.String(500, err.Error())
		}

		return c.JSON(200, listMaxCommision)

	}

}

func listarComercialesComision(cfg db.ApiConfig) echo.HandlerFunc {
	return func(c echo.Context) error {
		listCommercials, err := cfg.DB.AdCommission(context.Background())

		if err != nil {
			return c.String(500, err.Error())
		}

		return c.JSON(200, listCommercials)
	}

}

func listarPedidosAnio2017(cfg db.ApiConfig) echo.HandlerFunc {
	return func(c echo.Context) error {
		listProducts, err := cfg.DB.ListProducts2007(context.Background())

		if err != nil {
			return c.String(500, err.Error())
		}

		return c.JSON(200, listProducts)

	}

}

func listarIDsClientes(cfg db.ApiConfig) echo.HandlerFunc {
	return func(c echo.Context) error {
		listClients, err := cfg.DB.ListIDClientMakePurchase(context.Background())

		if err != nil {
			return c.String(500, err.Error())
		}

		return c.JSON(200, listClients)

	}

}

func listarPedidos(cfg db.ApiConfig) echo.HandlerFunc {
	return func(c echo.Context) error {

		listProducts, err := cfg.DB.ListProducts(context.Background())

		if err != nil {
			return c.String(500, err.Error())
		}

		return c.JSON(200, listProducts)

	}

}

func listarDosPedidosMayorValor(cfg db.ApiConfig) echo.HandlerFunc {

	return func(c echo.Context) error {
		twoOrders, err := cfg.DB.ListProductsLimitTwo(context.Background())

		if err != nil {
			return c.String(500, err.Error())
		}

		return c.JSON(200, twoOrders)

	}

}
