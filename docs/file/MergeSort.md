## 前言

> ​     「 归并排序 」是利用了 「分而治之 」的思想，进行递归计算的排序算法，效率在众多排序算法中的佼佼者，一般也会出现在各种「数据结构」的教科书上。



## 一、简单释义



### 1、算法目的

​        将原本乱序的数组变成有序，可以是「升序」或者「降序」（为了描述统一，本文一律只讨论「 升序」的情况）。



### 2、算法思想

​       通过将当前乱序数组分成长度近似的两份，分别进行「 递归 」调用，然后再对这两个排好序的数组，利用两个游标，将数据元素依次比较，选择相对较小的元素存到一个「 辅助数组 」中，再将「 辅助数组 」中的数据存回「 原数组 」。



### 3、命名由来

​        每次都是将数列分成两份，「 递归 」分别排序后，再进行「 合并 」</b></font>，故此命名「 归并排序 」。





## 二、核心思想

「递归」：函数通过改变参数，自己调用自己。

「比较」：关系运算符 小于(<) 的运用。

「合并」：两个数组合并成一个数组的过程。





## 三、动图演示

### 1、样例

![img](./MergeSort.assets/FksRDNwpqp2X8AQXbJiXHtb05eRH.png)

​       初始情况下的数据如图所示，基本属于乱序，纯随机出来的数据。



![img](./MergeSort.assets/FtevGCrUiFGXAOhMajxwOjilrH8N.png)

###  

### 2、算法演示



## 四、算法前置



### 1、递归的实现

​        这个算法本身需要做一些「 递归 」计算，所以你至少需要知道「 递归 」的含义，这里以「 C语言 」为例，来看下一个简单的「 递归 」是怎么写的。代码如下：



```c
int sum(int n) {
    if(n <= 0) {
        return 0;
    } 
    return sum(n - 1) + n;
}
```





这就是一个经典的递归函数，求的是从 1 到 n 的和，那么我们把它想象成 1 到 n-1 的和再加 n，而  1 到 n-1 的和为 sum(n-1)，所以整个函数体就是两者之和，这里 sum(n) 调用 sum(n-1) 的过程就被称为「 递归 」。



### 2、比较的实现

​      「比较」两个元素的大小，可以采用关系运算符，本文我们需要排序的数组是按照「升序」排列的，所以用到的关系运算符是「小于运算符（即 <）」。

​       我们可以将两个数的「比较」写成一个函数 smallerThan，以「 C语言 」为例，实现如下



```c
#define Type int
bool smallerThan(Type a, Type b) {
    return a < b;
}
```





其中 Type 代表数组元素的类型，可以是整数，也可以是浮点数，甚至可以是一个类，这里我们统一用 int 来讲解，即 32位有符号整型。



### 3、合并的实现

​        所谓「合并」，就是将两个有序数组合成一个有序数组的过程。

​        如下图所示：「 红色数组 」和 「 黄色数组 」各自有序，然后通过一个额外的数组，「合并」计算后得到一个有序的数组。





## 五、算法描述



### 1、问题描述

> ​        给定一个 n 个元素的数组，数组下标从 0 开始，采用「 归并排序 」将数组按照「升序」排列。



### 2、算法过程

> ​       整个算法的执行过程用 mergeSort(a[], l, r) 描述，代表 当前待排序数组 a，左区间下标 l，右区间下标 r，分以下几步：
>
> ​       1） 计算中点 mid = (l + r) / 2；
>
> ​       2）递归调用 mergeSort(a[], l, mid) 和 mergeSort(a[], mid+1, r)；
>
> ​       3）将 2）中两个有序数组进行有序合并，再存储到 a[l:r]；
>
> ​       4）实际调用时，调用 mergeSort(a[], 0, n-1) 就能得到整个数组的排序结果。



​       「 递归 」是自顶向下的，实际上程序真正运行过程是自底向上「 回溯 」的过程：给定一个 n 个元素的数组，「 归并排序 」将执行如下几步：

​         1）将每对单个元素归并为 2个元素 的有序数组；

​         2）将 2个元素 的每对有序数组归并成 4个元素 的有序数组，重复这个过程…；

​         3）最后一步：合并 2 个 n / 2 个元素的排序数组（为了简化讨论，假设 n 是偶数）以获得完全排序的 n 个元素数组。



## 六、算法分析



### 1、时间复杂度

​       我们假设 「比较」和「赋值」的时间复杂度为 O(1)。

​       我们首先讨论「 归并排序 」算法的最重要的子程序：O(n) 归并，然后解析这个归并排序算法。

​       给定两个大小为 n1 和 n2 的排序数组 A 和 B，我们可以在 O(n) 时间内将它们有效地归并成一个大小为 n = n1 + n2 的组合排序数组。可以通过简单地比较两个数组的前面并始终取两个中较小的一个来实现的。 

​       问题是这个归并过程被调用了多少次？

​      由于每次都是对半切，所以整个归并过程类似于一颗二叉树的构建过程，次数就是二叉树的高度，即 log_n，所以归并排序的时间复杂度为 O(nlogn)。



### 2、空间复杂度

​       由于归并排序在归并过程中需要额外的一个「 辅助数组 」，并且最大长度为原数组长度，所以「 归并排序 」的空间复杂度为 O(n)。





## 七、优化方案

​      「 归并排序 」在众多排序算法中效率较高，时间复杂度为 O(nlogn) 。 

​        但是，由于归并排序在归并过程中需要额外的一个「 辅助数组 」，所以申请「 辅助数组 」内存空间带来的时间消耗会比较大，比较好的做法是，实现用一个和给定元素个数一样大的数组，作为函数传参传进去，所有的 「 辅助数组 」干的事情，都可以在这个传参进去的数组上进行操作，这样就免去了内存的频繁申请和释放。



## 八、源码详解

```c
#include <stdio.h>
#include <malloc.h>
 
#define maxn 1000001

int a[maxn];

void Input(int n, int *a) {
    for(int i = 0; i < n; ++i) {
        scanf("%d", &a[i]);
    }
}

void Output(int n, int *a) {
    for(int i = 0; i < n; ++i) {
        if(i)
            printf(" ");
        printf("%d", a[i]);
    }
    puts("");
}

void MergeSort(int *nums, int l, int r) {
    int i, mid, p, lp, rp;
    int *tmp = (int *)malloc( (r-l+1) * sizeof(int) );    // (1)  
    if(l >= r) {
        return ;                                          // (2) 
    }
    mid = (l + r) >> 1;                                   // (3) 
    MergeSort(nums, l, mid);                              // (4) 
    MergeSort(nums, mid+1, r);                            // (5) 
    p = 0;                                                // (6) 
    lp = l, rp = mid+1;                                   // (7) 
    while(lp <= mid || rp <= r) {                         // (8) 
        if(lp > mid) {
            tmp[p++] = nums[rp++];                        // (9) 
        }else if(rp > r) {
            tmp[p++] = nums[lp++];                        // (10) 
        }else {
            if(nums[lp] <= nums[rp]) {                    // (11) 
                tmp[p++] = nums[lp++];
            }else {
                tmp[p++] = nums[rp++];
            }
        }
    }
    for(i = 0; i < r-l+1; ++i) {
        nums[l+i] = tmp[i];                               // (12) 
    } 
    free(tmp);                                            // (13) 
}

int main() {
    int n;
    while(scanf("%d", &n) != EOF) {
        Input(n, a);
        MergeSort(a, 0, n-1);
        Output(n, a);
    }
    return 0;
} 


(1) 申请一个辅助数组，用于对原数组进行归并计算；

(2) 只有一个元素，或者没有元素的情况，则不需要排序；

(3) 将数组分为 [l, mid] 和 [mid+1, r] 两部分；

(4) 递归排序 [l, mid] 部分；

(5) 递归排序 [mid+1, r] 部分；

(6) 将需要排序的数组缓存到 tmp 中，用 p 作为游标；

(7) 初始化两个数组的指针；

(8) 当两个指针都没有到结尾，则继续迭代；

(9) 只剩下右边的数组，直接排；

(10) 只剩下走右边的数组，直接排；

(11) 取小的那个先进 tmp 数组；

(12) 别忘了将排序好的数据拷贝回原数组；

(13) 别忘了释放临时数据，否则就内存泄漏了！！！
```





### Java版

```java
package Java.DataStructuresAndAlgorithms.Algorithms.sort;

import java.text.SimpleDateFormat;
import java.util.Date;

public class MergeSort {

    public static void main(String[] args) {

        int[] arr = new int[8000000];
        for (int i = 0; i < 8000000; i++) {
            arr[i] = (int) (Math.random() * 8000000);
        }
        Date date1 = new Date();
        SimpleDateFormat simpleDateFormat = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
        String date1Str = simpleDateFormat.format(date1);
        System.out.println("排序前的时间是=" + date1Str);

        int temp[] = new int[arr.length];
        mergeSort(arr, 0, arr.length - 1, temp);

        Date date2 = new Date();
        // SimpleDateFormat simpleDateFormat = new SimpleDateFormat("yyyy-MM-dd
        // HH:mm:ss");
        String date2Str = simpleDateFormat.format(date2);
        System.out.println("排序后的时间是=" + date2Str);
    }

    // 分+合的方法
    public static void mergeSort(int[] arr, int left, int right, int[] temp) {
        if (left < right) {
            int mid = (left + right) / 2;
            // 左递归分解
            mergeSort(arr, left, mid, temp);
            // 左递归分解
            mergeSort(arr, mid + 1, right, temp);
            // 合并
            merge(arr, left, mid, right, temp);
        }
    }

    // 合并的方法
    /**
     * 
     * @param arr   排序的原始的数组
     * @param left  索引
     * @param mid
     * @param right
     * @param temp  中转的数组
     */
    public static void merge(int[] arr, int left, int mid, int right, int[] temp) {
        int i = left; // 初始化i，左边有序序列的初始索引
        int j = mid + 1; // 初始化j，右边的索引
        int t = 0; // 指向temp数组的当前索引

        while (i <= mid && j <= right) {
            // 如果左边的有序序列的当前的元素，是小于等于这个右边的
            // 那么我们将左边的元素填充到temp的数组之中
            // 然后t++，i++
            if (arr[i] <= arr[j]) {
                temp[t] = arr[i];
                t += 1;
                i += 1;
            } else {
                temp[t] = arr[j];
                t += 1;
                j += 1;
            }
        }
        // 把有剩余数据的一边全部都放到这个temp这个数组当中去
        while (i <= mid) { // 左边的元素有序序列还有剩余的元素，全部进行一个填充
            temp[t] = arr[i];
            t += 1;
            i += 1;
        }
        while (j <= right) {
            temp[t] = arr[j];
            t += 1;
            j += 1;
        }

        // 将temp这个数组中的元素copy到这个arr中
        // 注意这个并不是每次都是全部的拷贝所有

        t = 0;
        int tempLeft = left;
        while (tempLeft <= right) {
            arr[tempLeft] = temp[t];
            t += 1;
            tempLeft += 1;
        }
    }

}

```

