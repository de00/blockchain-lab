# LAB3实验报告

## **【实验名称】**Fabric搭建peer并加入通道

**姓名 李泓民 学号 PB18071495**

## **【实验过程】**

### step1

> 使用CA服务器注册身份，并获得CA服务器颁发的身份证书

在助教给出的kecheng账号scp相应的文件夹，对于给出的命令修改相应的参数，执行`fabric-ca-client register --id.name leehm --id.secret 123456 --id.type peer -u http://222.195.70.186:7054 --mspdir ca-msp​`,注册节点:

![image-20210610162334598](E:\OneDrive - mail.ustc.edu.cn\Files\learningMaterials\Courses2021Spring\区块链\labs\reports\lab3.assets\image-20210610162334598.png)

然后再把注册成功之后的peer的mspenroll到本地:

![image-20210617120621348](E:\OneDrive - mail.ustc.edu.cn\Files\learningMaterials\Courses2021Spring\区块链\labs\reports\lab3.assets\image-20210617120621348.png)

查看自己的证书:

![image-20210617120540168](E:\OneDrive - mail.ustc.edu.cn\Files\learningMaterials\Courses2021Spring\区块链\labs\reports\lab3.assets\image-20210617120540168.png)

### step2

> 在本地准备Peer节点启动所需要的文件，启动Peer节点

查看config.yaml文件:

![image-20210610165528493](E:\OneDrive - mail.ustc.edu.cn\Files\learningMaterials\Courses2021Spring\区块链\labs\reports\lab3.assets\image-20210610165528493.png)

使用

```
export FABRIC_CFG_PATH=/home/UserPB18071495/peer
```

指定上传的core.yaml的路径,然后修改 peer 节点的 config.yaml 配置文件,修改参数:

![image-20210617122541439](E:\OneDrive - mail.ustc.edu.cn\Files\learningMaterials\Courses2021Spring\区块链\labs\reports\lab3.assets\image-20210617122541439.png)

根据cacerts下面的文件名修改为222-195-70-186-7054.pem.

然后是修改core.yaml文件:

- 查询到9999端口是可以使用的,指派9999端口给listenAddress.node88代表的含义和222.195.70.188是一样的.

![image-20210617215329002](E:\OneDrive - mail.ustc.edu.cn\Files\learningMaterials\Courses2021Spring\区块链\labs\reports\lab3.assets\image-20210617215329002.png)

![image-20210617215908127](E:\OneDrive - mail.ustc.edu.cn\Files\learningMaterials\Courses2021Spring\区块链\labs\reports\lab3.assets\image-20210617215908127.png)

修改存储数据的路径:

![image-20210617220546595](E:\OneDrive - mail.ustc.edu.cn\Files\learningMaterials\Courses2021Spring\区块链\labs\reports\lab3.assets\image-20210617220546595.png)

将mspConfigPath设置为拷贝的助教的具有admin权限的msp:

![image-20210617220636247](E:\OneDrive - mail.ustc.edu.cn\Files\learningMaterials\Courses2021Spring\区块链\labs\reports\lab3.assets\image-20210617220636247.png)

将chaincodeListenAddress设置为:

![image-20210617220839269](E:\OneDrive - mail.ustc.edu.cn\Files\learningMaterials\Courses2021Spring\区块链\labs\reports\lab3.assets\image-20210617220839269.png)

修改snapshots储存路径:

![image-20210617221300855](E:\OneDrive - mail.ustc.edu.cn\Files\learningMaterials\Courses2021Spring\区块链\labs\reports\lab3.assets\image-20210617221300855.png)

设置完毕后启动节点:

![image-20210617124506003](E:\OneDrive - mail.ustc.edu.cn\Files\learningMaterials\Courses2021Spring\区块链\labs\reports\lab3.assets\image-20210617124506003.png)

获取配置区块

```
peer channel fetch config bcclass.block -c bcclass --orderer 222.195.70.186:7050
```

peer加入通道（加入通道的创世区块可以使用peer channel fetch获得）

```
peer channel join -b bcclass.block
```

![image-20210617180955596](E:\OneDrive - mail.ustc.edu.cn\Files\learningMaterials\Courses2021Spring\区块链\labs\reports\lab3.assets\image-20210617180955596.png)

最后调用peer channel list查看已加入的节点:

![image-20210617222913777](E:\OneDrive - mail.ustc.edu.cn\Files\learningMaterials\Courses2021Spring\区块链\labs\reports\lab3.assets\image-20210617222913777.png)

