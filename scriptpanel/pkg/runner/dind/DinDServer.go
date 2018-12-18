package scriptpanel

import (
	"ScriptPanel/scriptpanel/pkg/objects"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	dockerClient "github.com/docker/docker/client"
)

const (
	//dindURL             = "http://10.10.10.12:4444"
	dockerClientVersion = "v1.37"
)

//DinD type is the driver object that talks with the DinD server
type DinD struct {
	*dockerClient.Client
}

func getURL() (ip, port string) {
	return os.Getenv("DIND_IP"), os.Getenv("DIND_PORT")
}

func (dind *DinD) downloadImage(image string) {
	ctx := context.Background()

	_, err := dind.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot download the image: %v", err)
	}

	//defer out.Close()
	//io.Copy(os.Stdout, out)
}

// DownloadImages orders the DinD to pull the images
func (dind *DinD) DownloadImages(images []string) {
	for _, image := range images {
		go dind.downloadImage(image)
	}
}

// NewDinDServer returns a configured DinD object
func NewDinDServer() *DinD {
	ip, port := getURL()

	dindURL := fmt.Sprintf("http://%s:%s", ip, port)

	defaultHeaders := map[string]string{"User-Agent": "engine-api-cli-1.0"}

	netTransport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5000 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5000 * time.Second,
	}

	httpClient := &http.Client{
		Timeout:   time.Second * 1000,
		Transport: netTransport,
	}

	client, dindConnectionError := dockerClient.NewClient(dindURL, dockerClientVersion, httpClient, defaultHeaders)

	if dindConnectionError != nil {
		fmt.Fprintf(os.Stderr, "Error getting a connection with dind: %v", dindConnectionError)
		return nil
	}

	images := []string{
		"bash:4.4",
		"python:2.7",
		"perl:5.24",
		"ruby:2.5",
	}

	dind := &DinD{
		client,
	}

	dind.DownloadImages(images)

	return dind
}

//RunScript runs a script in a DinD container that is stored previosuly
func (dind *DinD) RunScript(runRequest objects.RunRequest) string {
	ctx := context.Background()

	command := make([]string, 0)

	if runRequest.ScriptObject.Language != "rb" { //ruby doesn't work the same way as the others containers do
		command = append(command, langHash[runRequest.ScriptObject.Language])
	}

	command = append(command, "/home/"+runRequest.ScriptObject.Filename)
	command = append(command, runRequest.ConfigObject.Flag)

	for _, arg := range runRequest.ConfigObject.Args {
		command = append(command, arg)
	}

	resp, containerCreateError := dind.ContainerCreate(
		ctx,
		&container.Config{
			Image: imgHash[runRequest.ScriptObject.Language] + ":" + verHash[runRequest.ScriptObject.Language],
			Cmd:   command,
			Tty:   true,
		},
		&container.HostConfig{
			Mounts: []mount.Mount{
				{
					Type:   mount.TypeBind,
					Source: "/home/scripts/admin/" + runRequest.ScriptObject.Company + "/" + runRequest.ScriptObject.Department + "/" + runRequest.ScriptObject.Username + "/",
					Target: "/home/",
				},
			},
		}, nil, "")

	if containerCreateError != nil {
		log.Println(containerCreateError)
		fmt.Fprintf(os.Stderr, "Error creating the container: %v", containerCreateError)
		return "Cannot create the container"
	}

	if containerStartError := dind.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); containerStartError != nil {
		log.Println(containerStartError)
		fmt.Fprintf(os.Stderr, "Error starting the container: %v", containerStartError)
		return "Cannot start the container"
	}

	statusCh, errCh := dind.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)

	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	// status, err := dind.ContainerWait(ctx, resp.ID)
	// if err != nil || status != http.StatusOK {
	// 	panic(err)
	// }

	out, err := dind.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	result, err := ioutil.ReadAll(out)

	return string(result)
}

type consult map[string]string

var imgHash = consult{
	"py": "python",
	"sh": "bash",
	"rb": "ruby",
	"pl": "perl",
}

var langHash = consult{
	"py": "python",
	"sh": "bash",
	"rb": "",
	"pl": "perl",
}

var verHash = consult{
	"py": "2.7",
	"sh": "4.4",
	"rb": "2.5",
	"pl": "5.24",
}
