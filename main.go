package main

import (
	"fmt"
	"io/ioutil"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"log"
	"os"
)


/*
The application expects environment varialbe "KUBECONFIG" to be set, then uninstalls Service Catalog and removes all SC resources.
 */
func main() {
	// read the kubeconfig
	kcContent, err := ioutil.ReadFile(os.Getenv("KUBECONFIG"))
	if err != nil {
		panic(err)
	}

	cleaner, err := NewCleaner(kcContent)
	if err != nil {
		panic(err)
	}

	log.Println("Removing Service CAtalog release")
	cleaner.RemoveServiceCatalogRelease()

	log.Println()
	log.Println("Removing finalizers")
	err = cleaner.PrepareForRemoval()
	fmt.Println(err)

	log.Println()
	log.Println("Deleting resources")
	err = cleaner.RemoveResources()
	fmt.Println(err)
}

/*
1. delete helm release of SC in namespace kyma-system
 - check if SC exists
 - delete the release
2. find all ServiceBindings
 - remove owner references/finalizers
 - delete all SB
3. find all (cluster)servicebrokers
 - remove finalizers
 - delete all (c)SB
4. find all serviceinstances
 - remove finalizers
 - delete them
*/
