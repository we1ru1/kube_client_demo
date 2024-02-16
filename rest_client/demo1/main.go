package main

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/internalversion/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

/*
*

	获取kube-system命名空间下的pod列表
*/
func main() {
	/*
		1. 拿到k8s配置文件
		2.保证开发机器能够通过配置文件连接到集群
	*/

	// 1. 拿到k8s配置文件，生成config对象
	config, err := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	if err != nil {
		panic(err)
	}

	// 2.配置API路径
	config.APIPath = "api" //  api是核心
	//config.APIPath = "apis" // apis是扩展

	// 3.配置分组版本
	config.GroupVersion = &corev1.SchemeGroupVersion //无组名资源组

	// 4.配置数据的编解码工具
	config.NegotiatedSerializer = scheme.Codecs

	// 5.实例化RESTClient对象
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err.Error())
	}

	// 6.调用RESTClient对象的Get方法，获取pod列表
	result := corev1.PodList{}

	// 跟apiserver交互
	err = restClient.
		Get().                                                         // Get请求方式
		Namespace("kube-system").                                      // 命名空间
		Resource("pods").                                              // 资源名称
		VersionedParams(&metav1.ListOptions{}, scheme.ParameterCodec). // 参数及参数的序列化工具
		Do(context.TODO()).                                            // 发送请求
		Into(&result)                                                  // 把返回结果写道result中
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(result)
}
