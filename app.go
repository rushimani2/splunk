package main

import (
    "example.com/cdk8s-demo/imports/k8s"
    "github.com/aws/constructs-go/constructs/v10"
    "github.com/aws/jsii-runtime-go"
)

type AppProps struct {
    Image         *string
    Replicas      *float64
    Port          *float64
    ContainerPort *float64
    Name          *string
}

func NewDeployment(scope constructs.Construct, id *string, props *AppProps) constructs.Construct {
    construct := constructs.NewConstruct(scope, id)

    replicas := props.Replicas
    if replicas == nil {
        replicas = jsii.Number(1)
    }

    containerPort := props.ContainerPort
    if containerPort == nil {
        containerPort = jsii.Number(8080)
    }

    appName := props.Name
    if appName == nil {
        appName = jsii.String("go-web-app")
    }

    label := map[string]*string{
        "app": appName,
    }

    k8s.NewKubeDeployment(construct, jsii.String("deployment"), &k8s.KubeDeploymentProps{
        Metadata: &k8s.ObjectMeta{
            Name:   appName,
            Labels: &label,
        },
        Spec: &k8s.DeploymentSpec{
            Replicas: replicas,
            Selector: &k8s.LabelSelector{MatchLabels: &label},
            Template: &k8s.PodTemplateSpec{
                Metadata: &k8s.ObjectMeta{Labels: &label},
                Spec: &k8s.PodSpec{
                    Containers: &[]*k8s.Container{{
                        Name:  jsii.String("web"),
                        Image: props.Image,
                        Ports: &[]*k8s.ContainerPort{{
                            ContainerPort: containerPort,
                        }},
                    }},
                },
            },
        },
    })

    return construct
}

func NewService(scope constructs.Construct, id *string, props *AppProps) constructs.Construct {
    construct := constructs.NewConstruct(scope, id)

    port := props.Port
    if port == nil {
        port = jsii.Number(80)
    }

    containerPort := props.ContainerPort
    if containerPort == nil {
        containerPort = jsii.Number(8080)
    }

    appName := props.Name
    if appName == nil {
        appName = jsii.String("go-web-app")
    }

    label := map[string]*string{
        "app": appName,
    }

    k8s.NewKubeService(construct, jsii.String("service"), &k8s.KubeServiceProps{
        Metadata: &k8s.ObjectMeta{
            Name:   jsii.String(*appName + "-service"),
            Labels: &label,
        },
        Spec: &k8s.ServiceSpec{
            Type: jsii.String("LoadBalancer"),
            Ports: &[]*k8s.ServicePort{{
                Port:       port,
                TargetPort: k8s.IntOrString_FromNumber(containerPort),
            }},
            Selector: &label,
        },
    })

    return construct
}
