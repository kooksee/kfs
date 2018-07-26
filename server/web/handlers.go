package web

import (
	"github.com/labstack/echo"
	"github.com/kooksee/kfs/types"
	"net/http"
)

func index(c echo.Context) error {
	req := types.RPCRequest{}
	if err := c.Bind(req); err != nil {
		return err
	}

	switch req.Method {
	default:
		return c.JSON(http.StatusNotFound, types.RPCResponse{Code: 404, Msg: "方法不存在"})

	case "file.add":
	case "file.list":
	case "file.rm":
	case "file.get":
	case "file.pin":

	case "metadata.ls":
	case "metadata.update":
	case "metadata.add":
	case "metadata.rm":

	case "peer.ls":
	case "peer.rm":
	case "peer.add":
	}

	return nil
}
