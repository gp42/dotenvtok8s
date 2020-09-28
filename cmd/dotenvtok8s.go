package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"path"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/gp42/dotenvtok8s/pkg/util"
)

var (
	version	string
	flagInFile, flagOutDir, flagSecretPrefix string
	flagVersion bool
)


func main() {
	flag.BoolVar(&flagVersion, "version", false, "Show version and exit")
	flag.StringVar(&flagInFile, "in-file", ".env", "Specify .env file")
	flag.StringVar(&flagOutDir, "out-dir", "", "Specify output dir")
	flag.StringVar(&flagSecretPrefix, "secret-prefix", "", "Specify a prefix to generate secrets instead (prefix will be removed from generated env vars, repeat prefix twice to have it)")
	flag.Parse()

	if flagVersion {
		fmt.Printf("Version: %s\n", version)
		return
	}

	data, err := godotenv.Read(flagInFile)
	util.Check(err)

	cmData, secretData := util.SplitKeys(data, flagSecretPrefix)

	cm := &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{Name: "env-cm"},
		Data: cmData,
	}

	err = util.WriteYaml(cm, path.Join(flagOutDir, "cm.yaml"))
	util.Check(err)

	if len(secretData) > 0 {
		secret := &corev1.Secret{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Secret",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{Name: "env-secret"},
			StringData: secretData,
		}

		err = util.WriteYaml(secret, path.Join(flagOutDir, "secret.yaml"))
		util.Check(err)
	}
}
