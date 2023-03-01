### `kubernetes`

##### 了解

- [发展](https://zhuanlan.zhihu.com/p/266180037)

- [架构](https://feisky.gitbooks.io/kubernetes/content/architecture/architecture.html)

- [组件](https://kubernetes.io/zh-cn/docs/concepts/overview/components/)

##### 环境搭建

- `minikube`

    - [minikube start | minikube (k8s.io)](https://minikube.sigs.k8s.io/docs/start/)

      ```shell
      # 启动集群
      minikube start
      # 查看节点。kubectl 是一个用来跟 K8S 集群进行交互的命令行工具
      kubectl get node
      # 停止集群
      minikube stop
      # 清空集群
      minikube delete --all
      # 安装集群可视化 Web UI 控制台
      minikube dashboard
      ```

    - [手册](https://minikube.sigs.k8s.io/docs/handbook/)

- `docker`

    - 设置里勾选一下添加一下重启

      ![image-20230301184352743](https://cdn.jsdelivr.net/gh/frenude/images@main/2023/03/01/8fcb2622aae63c6c96fb31c808c2daa6.png)

      、

- 会起一大堆docker，看着墨迹，还是单节点，没意义。

- `kubeadm`

    - 适合生产裸机搭建
    - [提供一个搭建文档](https://k8s.easydoc.net/docs/dRiQjyTY/28366845/6GiNOzyZ/nd7yOvdY)
    - 网上视频教程全是

- 云平台

##### `Pod`

-  ReplicationController（RC）：RC保证了在所有时间内，都有特定数量的Pod副本正在运行，如果太多了，RC就杀死几个，如果太少了，RC会新建几个
-  **ReplicaSet（RS）**：代用户创建指定数量的pod副本数量，确保pod副本数量符合预期状态，并且支持滚动式自动扩容和缩容功能。
-  **Deployment**（重要）：工作在ReplicaSet之上，用于管理无状态应用，目前来说最好的控制器。支持滚动更新和回滚功能，还提供声明式配置。
-  **DaemonSet**：用于确保集群中的每一个节点只运行特定的pod副本，通常用于实现系统级后台任务。比如ELK服务
-  Job：只要完成就立即退出，不需要重启或重建。
-  CronJob：周期性任务控制，不需要持续后台运行
-  **StatefulSet**：管理有状态应用。需要数据持久化的服务mysql等

##### `kubectl`

- **常用指令**

  ```shell
  # 部署应用
  kubectl apply -f app.yaml
  # 查看 deployment
  kubectl get deployment
  # 查看 pod
  kubectl get pod -o wide
  # 查看 pod 详情
  kubectl describe pod pod-name
  # 查看 log
  kubectl logs pod-name
  # 进入 Pod 容器终端， -c container-name 可以指定进入哪个容器。
  kubectl exec -it pod-name -- bash
  # 伸缩扩展副本
  kubectl scale deployment test-k8s --replicas=5
  # 把集群内端口映射到节点
  kubectl port-forward pod-name 8090:8080
  # 查看历史
  kubectl rollout history deployment test-k8s
  # 回到上个版本
  kubectl rollout undo deployment test-k8s
  # 回到指定版本
  kubectl rollout undo deployment test-k8s --to-revision=2
  # 删除部署
  kubectl delete deployment test-k8s
  # 查看全部
  kubectl get all
  # 重新部署
  kubectl rollout restart deployment test-k8s
  # 命令修改镜像，--record 表示把这个命令记录到操作历史中
  kubectl set image deployment test-k8s test-k8s=kicbase/echo-server:1.0 -with-error --record
  # 暂停运行，暂停后，对 deployment 的修改不会立刻生效，恢复后才应用设置
  kubectl rollout pause deployment test-k8s
  # 恢复
  kubectl rollout resume deployment test-k8s
  # 输出到文件
  kubectl get deployment test-k8s -o yaml >> app2.yaml
  # 删除全部资源
  kubectl delete all --all
  ```

- 创建`pod`

    - `kubectl run k8s-hello --image=kicbase/echo-server:1.0`
    - `kubectl run k8s-hello --image=kicbase/echo-server:1.0  --dry-run` 这个是指令目前最有用的

  ```yaml
  apiVersion: v1
  kind: Pod
  metadata:
    name: k8s-hello
  spec:
    # 定义容器，可以多个
    containers:
      - name: k8s-hello # 容器名字
        image: kicbase/echo-server:1.0 # 镜像
  
  ```

- **创建`deployment`**

    - `kubectl create deployment hello-k8s --image=kicbase/echo-server:1.0`

      ```yaml
      apiVersion: apps/v1
      kind: Deployment
      metadata:
        # 部署名字
        name: hello-k8s
      spec:
        replicas: 2
        # 用来查找关联的 Pod，所有标签都匹配才行
        selector:
          matchLabels:
            app: hello-k8s
        # 定义 Pod 相关数据
        template:
          metadata:
            labels:
              app: hello-k8s
          spec:
            # 定义容器，可以多个
            containers:
            - name: hello-k8s # 容器名字
              image: kicbase/echo-server:1.0 # 镜像
      
      ```

- 剩下的去看官网吧[工作负载 | Kubernetes](https://kubernetes.io/zh-cn/docs/concepts/workloads/)

