package api

import (
	"context"
	"time"

	goHttp "net/http"

	"github.com/pterm/pterm"
	httpIface "github.com/taubyte/http"
	http "github.com/taubyte/http/basic"
	"github.com/taubyte/http/options"
	"github.com/taubyte/tau/libdream/common"
	"github.com/taubyte/tau/libdream/services"
)

type multiverseService struct {
	rest httpIface.Service
	common.Multiverse
}

func BigBang() error {
	var err error

	srv := &multiverseService{
		Multiverse: services.NewMultiVerse(),
	}

	srv.rest, err = http.New(srv.Context(), options.Listen(common.DreamlandApiListen), options.AllowedOrigins(true, []string{".*"}))
	if err != nil {
		return err
	}

	srv.setUpHttpRoutes().Start()

	waitCtx, waitCtxC := context.WithTimeout(srv.Context(), 10*time.Second)
	defer waitCtxC()

	for {
		select {
		case <-waitCtx.Done():
			return waitCtx.Err()
		case <-time.After(100 * time.Millisecond):
			if srv.rest.Error() != nil {
				pterm.Error.Println("Dreamland failed to start")
				return srv.rest.Error()
			}
			_, err := goHttp.Get("http://" + common.DreamlandApiListen)
			if err == nil {
				pterm.Info.Println("Dreamland ready")
				return nil
			}
		}
	}
}

/* Comments by Aarshee Bhattacharya
I read the code and understood that it is from a Go application. It defined a multiverseService and a BigBang function and as far as my 
understanding goes it created an HTTP server using 'github.com/taubyte/package' and it is waiting until it getrs ready.
Lets break down the code to understand it better-
1. The code imports multiple Go packages and even pterm for a colourful terminal output.
2. The multiverseService has two fields one is the rest tahts an HTTP service and common.multiverse that is a custom type from the libdream
package.
3. The BigBang() function is used to initialize some variables including an error variable. It creates an instance for the multiverseService
and sets up an HTTP server. The HTTP serevr then starts and routes are set using teh setUpHttpRoutes method of the multiverseService.
There is even a timeout that wont let it start unless 10 seconds have passed.
4. The Server Startup Loop is used to check the timeout context when it's done otherwise it waits for 10 milliseconds. It prints an error
message if there's a problem starting the HTTP server. It makes an HTTp request to a specified address and even prints a message when the 
action was suucessful.

Therefore, in simple words the code is setting up an HTTP serevr and waits for it to start and if it starts successfully it prints a message
saying the Dreamland ready otherwise Dreamland failed to start.
