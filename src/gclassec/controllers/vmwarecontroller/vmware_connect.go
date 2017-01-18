package vmwarecontroller

import (
	"context"
	"flag"
	"fmt"
	"net/url"
	"os"
	"text/tabwriter"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/property"
	//"github.com/vmware/govmomi/units"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/units"
	"net/http"
	"encoding/json"
)

const (
	envURL = "https://110.110.110.140:443/sdk"
	envUserName = "administrator@vsphere.local"
	envPassword = "Vcenter#1234"
	envInsecure = "true"
)

type (

    UserController struct{}
)
func NewUserController() *UserController {
    return &UserController{}
}

type obj struct {
	VMName string
	OverallCPU int32
	GuestMemory int32
	StorageCommitted float32
}
var urlDescription = fmt.Sprintf("ESX or vCenter URL [%s]", envURL)
var urlFlag = flag.String("url", envURL, urlDescription)

var insecureDescription = fmt.Sprintf("Don't verify the server's certificate chain [%s]", envInsecure)
var insecureFlag = flag.Bool("insecure", true, insecureDescription)

func processOverride(u *url.URL) {
	//envUsername := os.Getenv(envUserName)
	//envPassword := os.Getenv(envPassword)

	// Override username if provided
	if envUserName != "" {
		var password string
		var ok bool

		if u.User != nil {
			password, ok = u.User.Password()
		}

		if ok {
			u.User = url.UserPassword(envUserName, password)
		} else {
			u.User = url.User(envUserName)
		}
	}

	// Override password if provided
	if envPassword != "" {
		var username string

		if u.User != nil {
			username = u.User.Username()
		}

		u.User = url.UserPassword(username, envPassword)
	}
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	os.Exit(1)
}

func   (uc UserController) GetDynamicVcenterDetails(w http.ResponseWriter, r *http.Request)(){
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	fmt.Println(*insecureFlag)

	flag.Parse()

	// Parse URL from string
	u, err := url.Parse(*urlFlag)
	if err != nil {
		exit(err)
	}

	// Override username and/or password as required
	processOverride(u)

	// Connect and log in to ESX or vCenter
	c, err := govmomi.NewClient(ctx, u, *insecureFlag)
	if err != nil {
		exit(err)
	}

	f := find.NewFinder(c.Client, true)

	// Find one and only datacenter
	dc, err := f.DefaultDatacenter(ctx)
	if err != nil {
		exit(err)
	}

	// Make future calls local to this datacenter
	f.SetDatacenter(dc)

	// Find virtual machines in datacenter
	vms, err := f.VirtualMachineList(ctx, "*")
	fmt.Println(vms)

	pc := property.DefaultCollector(c.Client)

	var refv []types.ManagedObjectReference
	for _, ds := range vms {
		refv = append(refv, ds.Reference())
	}

	// Retrieve name property for all vms
	var vmt []mo.VirtualMachine
	err = pc.Retrieve(ctx, refv, []string{"summary"}, &vmt)
	if err != nil {
  		exit(err)
	}

	// Print summary
	tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)

	fmt.Println("Virtual machines found:", len(vmt))
	for _, vm := range vmt {
  		//fmt.Fprintf(tw, "%s\n", vm.Name)
		fmt.Println("VM Name : ", vm.Summary.Config.Name)
		fmt.Println("Overall CPU : ", vm.Summary.QuickStats.OverallCpuUsage)
		fmt.Println("Guest memory : ", vm.Summary.QuickStats.GuestMemoryUsage)
		fmt.Println("Committed storage : ", units.ByteSize(vm.Summary.Storage.Committed))
		//_ = json.NewEncoder(os.Stdout).Encode(&vm)
		output := obj{vm.Summary.Config.Name,vm.Summary.QuickStats.OverallCpuUsage,vm.Summary.QuickStats.GuestMemoryUsage,float32(vm.Summary.Storage.Committed)/float32(1024*1024*1024)}
		_ = json.NewEncoder(w).Encode(output)
	}

	tw.Flush()
}
