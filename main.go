package main
import (
	"github.com/spf13/cobra"
	"os"
	"log"
	"text/template"
)


func main() {
	rootCmd := &cobra.Command{
		Use:   "Macro",
		Short: "CLI tool to create, run, and manage microservices",
		Long: `Create a new microservice with these options:
			- Service layer,
			- Data Access Layer,
			- Server Layer,
			- Workflow layer,
			- Deployment layer,
		`,}

	rootCmd.AddCommand(&cobra.Command{
		Use:   "create",
		Short: "Create a new microservice",
	})
	createService("go", "test")
	rootCmd.Execute()

}

func createService(lang string, name string) {
	generateService(name)
}

func generateService(name string) {
	
	path := "./" + name
	
	err := os.Mkdir(path, 0755)
	if err != nil {
        log.Fatalf("Failed to create directory: %v", err)
    }
	
	// Create grpc server folder
	err = os.Mkdir(path + "/server", 0755)
	if err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}
	server_file, err := os.Create(path + "/server/server.go")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer server_file.Close()

	err = os.Mkdir(path + "/server/proto", 0755)
	if err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}

	// Create service folder
	err = os.Mkdir(path + "/service", 0755)
	if err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}
	service_file, err := os.Create(path + "/service/service.go")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer service_file.Close()

	// Create data access layer folder
	err = os.Mkdir(path + "/store", 0755)
	if err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}

	store_file, err := os.Create(path + "/store/store.go")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer store_file.Close()


	// Create data access layer folder
	err = os.Mkdir(path + "/deploy", 0755)
	if err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}

	deploy_file, err := os.Create(path + "/deploy/deploy.yaml")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer deploy_file.Close()

	service_acc_file, err := os.Create(path + "/deploy/service_account.yaml")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	// Parse the serviceAccountTemplate
	tmpl, err := template.New("serviceAccount").Parse(serviceAccountTemplate)
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}

	// Define the data for the template
	data := ServiceAccount{
		Name:      "my-service-account",
		Namespace: "default",
	}

	// Execute the template and write the output to the file
	err = tmpl.Execute(service_acc_file, data)
	if err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}
	 


}


type ServiceAccount struct {
    Name      string
    Namespace string
}

const serviceAccountTemplate = `
apiVersion: v1
kind: ServiceAccount
metadata:
  name: {{.Name}}
  namespace: {{.Namespace}}
`