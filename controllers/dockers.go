package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

type Dockers struct {
	
}

type containerModel struct {
	ID string
	Name   string
	Status 	string
	Network string
}


func (d *Dockers) ListAll(ctx *gin.Context)  {
	dctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(dctx, types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
	}

	u := []containerModel{}

	for _, container := range containers {
		c, _ := cli.ContainerInspect(dctx, container.ID)
		a := containerModel{
			ID: container.ID,
			Name: container.Names[0],
			Status: container.Status,
			Network: c.NetworkSettings.IPAddress,
		}
		u = append(u, a)
				
	}
	
	ctx.JSON(http.StatusOK, gin.H{"dockers": u})
}


func (d *Dockers) StopContainer(ctx *gin.Context)  {
	id := ctx.Param("id")
	// fmt.Println(id)

	dctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	fmt.Print("Stopping container ",id, "... ")
	if err := cli.ContainerStop(dctx, id, nil); err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"dockers": "stop"})
}

func (d *Dockers) StartContainer(ctx *gin.Context)  {
	id := ctx.Param("id")
	// fmt.Println(id)

	dctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	fmt.Print("Starting container ",id, "... ")
	if err := cli.ContainerStart(dctx, id, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"dockers": "start"})
}