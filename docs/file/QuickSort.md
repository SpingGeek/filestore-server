#### 图解精讲快速排序（C语言实现）

「 快速排序 」是利用了「 分而治之 」的思想，进行递归计算的排序算法，效率在众多排序算法中的佼佼者，一般也会出现在各种「数据结构」的教科书上。于是，我也来简单讲一讲。



首先我们可能对于这种的分而治之的思想是很不习惯的，我们刚刚去接触这个快速排序这个算法的时候，我们可能是不太可以去真正的理解这种的思想的，而且我们在之前所学的比如说冒泡排序，插入排序，选择排序这样的相对简单的排序算法，这个快速排序我们是比较难入门的，但是没关系，听我细细道来！！！



### 一、简单释义

### 1、算法目的

​        将原本乱序的数组变成有序，可以是「升序」或者「降序」（为了描述统一，本文一律只讨论「 升序」的情况）。



### 2、算法思想

​       随机找到一个位置，将比它小的数都放到它「 左边 」，比它大的数都放到它「 右边 」，然后分别「 递归 」求解「 左边 」和「 右边 」使得两边分别有序。

![image-20240918201114567](./QuickSort.assets/image-20240918201114567.png)

![image-20240918201154082](./QuickSort.assets/image-20240918201154082.png)

### 3、命名由来

​       由于这个排序，相比其它的排序速度较快，故此命名「 快速排序 」。

#  

## 二、核心思想

「递归」：函数通过改变参数，自己调用自己。

「比较」：关系运算符 小于(<) 的运用。

「分治」：意为分而治之，先分，再治。将问题拆分成两个小问题，分别去解决。

## 三、动图演示



### 1、样例

​     初始情况下的数据如图所示，基本属于乱序，纯随机出来的数据。

![img](https://article-images.zsxq.com/FgHMj4WMNggnnaqcgTddFmq3NIKN)





### 2、算法演示

​       这个地方我们使用动图的形式进行展示也是比较不方便的，所以这里我分步骤给大家进行相关的讲解

1.首先我们可以去在这个数组里面先去选择一个基数作为这个标准，当然我们去选择这个基数的时候，在本质上这个是可以是一个随机的选择，当然依照我的习惯，我喜欢去选择这个里面的中间的这个数，虽然说这个是没有区别的，这个基数的选择是具有这个随机性的

2.之后，我们要去利用这个分而治之的思想，将原本没有关联的相关的数据产生相对的关系，“我们要去秉持着一个原则：就是我们要把这个基准数作为这个分割值，在基数左边的数值都要比这个基数的数值要去小，而在这个基数右边的数值都要比这个基数的数值要去大”，这样才可以方便我们的进一步进行一个排序

3.在这个基数的左右的两边，我们就去使用了这个递归的思想进行进一步的排序

4.随着在这之间的不同的多个数值的位置进行了一个相应的确定，这样我们的快速排序也变的越来越清晰明了



* ##### 时间复杂度是比较低的O（nlogn）



![image-20240918164044681](/Users/spring/Library/Application Support/typora-user-images/image-20240918164044681.png)



这个地方因为markdown是不方便我们去使用这个动图去进行展示的

所以我给大家一个链接，带着大家一起去观看这个快速排序的过程～



https://blog.csdn.net/justidle/article/details/104203963?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522EB860F24-4ABD-48F5-B2CA-2E70A9FDC40A%2522%252C%2522scm%2522%253A%252220140713.130102334..%2522%257D&request_id=EB860F24-4ABD-48F5-B2CA-2E70A9FDC40A&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~top_positive~default-1-104203963-null-null.142^v100^control&utm_term=%E5%BF%AB%E9%80%9F%E6%8E%92%E5%BA%8F&spm=1018.2226.3001.4187



## 四、算法前置



### 1、递归的实现

​        这个算法本身需要做一些「 递归 」计算，所以你至少需要知道「 递归 」的含义，这里以「 C语言 」为例，来看下一个简单的「 递归 」是怎么写的。代码如下：



```c
int sum(int n){
	if(n<=0) return 0;
	return sum(n-1)+n;
}
```

![image-20240918170012626](/Users/spring/Library/Application Support/typora-user-images/image-20240918170012626.png)

这就是一个经典的递归函数，求的是从 1 到 n 的和，那么我们把它想象成 1 到 n-1 的和再加 n，而  1 到 n-1 的和为 sum(n-1)，所以整个函数体就是两者之和，这里 sum(n) 调用 sum(n-1) 的过程就被称为「 递归 」。



### 2、比较的实现

​      「比较」两个元素的大小，可以采用关系运算符，本文我们需要排序的数组是按照「升序」排列的，所以用到的关系运算符是「小于运算符（即 <）」。

​        我们可以将两个数的「比较」写成一个函数 smallerThan，以「 C语言 」为例，实现如下：



```c
#define Type int
bool smallerThan(Type a, Type b) {
    return a < b;
}
```



 其中 Type 代表数组元素的类型，可以是整数，也可以是浮点数，也可以是一个类的实例，这里我们统一用int 来讲解，即 32位有符号整型。





### 3、分治的实现

​       所谓「分治」，就是把一个复杂的问题分成两个（或更多的相同或相似的）「 子问题 」，再把子问题分成更小的「 子问题 」……，直到最后子问题可以简单的直接求解，原问题的解即「 子问题 」的解的合并。

​       对于「 快速排序 」来说，我们选择一个基准数，将小于它的数都放到左边，大于它的数都放到它的右边，这个过程其实就是天然隔离了 **左边的数** 和 **右边的数**，使得两边的数 “分开”，这样就可以分开治理了。



## 五、算法描述



### 1、问题描述

> ​        给定一个 n 个元素的数组，数组下标从 0 开始，采用「 快速排序 」将数组按照「升序」排列。



### 2、算法过程

> ​       整个算法的执行过程用 quickSort(a[], l, r) 描述，代表 当前待排序数组 a，左区间下标 l，右区间下标 r，分以下几步：
>
> ​       1） 随机生成基准点 pivox = Partition(l, r)；
>
> ​       2）递归调用 quickSort(a[], l, pivox - 1) 和 quickSort(a[], pivox +1, r)；
>
> ​       3）Partition(l, r) 返回一个基准点，并且保证基准点左边的数都比它小，右边的数都比它大；Partition(l, r)称为分区。



## 六、算法分析

### 1、时间复杂度

​       首先，我们分析跑一次分区的成本。

​       在实现分区 Partition(l, r) 中，只有一个 for 循环遍历 (r - l) 次。 由于 r 可以和 n-1 一样大，i 可以低至 0，所以分区的时间复杂度是 O(n)。

​       类似于归并排序分析，快速排序的时间复杂度取决于分区被调用的次数。

​       当数组已经按照升序排列时，快速排序将达到最坏时间复杂度。总共 n 次分区，分区的时间计算如下：(n-1) + ... + 2 + 1 = n*(n-1) / 2。

​       总的时间复杂度为：O(n^2)。

​       当分区总是将数组分成两个相等的一半时，就会发生快速排序的最佳情况，如归并排序。当发生这种情况时，递归的深度只有 O(logn)。总的时间复杂度为 O(nlogn)。



## 七、优化方案

> ​      「 快速排序 」在众多排序算法中效率较高，平均时间复杂度为 O(nlogn)。但当完全有序时，最坏时间复杂度达到最坏情况 O(n^2)。 
>
> ​       所以每次在选择基准数的时候，我们可以尝试用随机的方式选取，这就是「 随机快速排序 」。
>
> ​       想象一下在随机化版本的快速排序中，随机化数据透视选择，我们不会总是得到 0，1 和 n-1 这种非常差的分割。所以不会出现上文提到的问题。



## 八、源码详解



### 1、快速排序实现1



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

void Swap(int *a, int *b) {
    int tmp = *a;
    *a = *b;
    *b = tmp;
}

int Partition(int a[], int l, int r){
    int i, j, pivox; 
    int idx = l + rand() % (r - l + 1);        // (1) 
    pivox = a[idx];                            // (2) 
    Swap(&a[l], &a[idx]);                      // (3) 
    i = j = l + 1;                             // (4) 
                                               // 
    while( i <= r ) {                          // (5) 
        if(a[i] < pivox) {                     // (6) 
            Swap(&a[i], &a[j]);                
            ++j;                               
        }
        ++i;                                   // (7) 
    }
    Swap(&a[l], &a[j-1]);                      // (8) 
    return j-1;
}


//递归进行划分
void QuickSort(int a[], int l, int r){
    if(l < r){
        int mid = Partition(a, l, r);
        QuickSort(a, l, mid-1);
        QuickSort(a, mid+1, r);
    }
}

int main() {
    int n;
    while(scanf("%d", &n) != EOF) {
        Input(n, a);
        QuickSort(a, 0, n-1);
        Output(n, a);
    }
    return 0;
} 
```

(1) 随机选择一个基准；

(2) pivox 代表基准值；

(3) 将基准和最左边的值交换；

(4) i 和 j 是两个同步指针，都从 l+1 开始；j-1 以后的数都是大于等于 基准值 的；

(5) 开始遍历整个排序区间，i 一定比 j 走的快，当 i 到达最右边的位置时，遍历结束； 

(6) 如果比基准值小的，交换 a[i] 和 a[j]，并且自增 j；

(7) 每次遍历 i 都需要自增；

(8) 第 j 个元素以后都是不比基准值小的元素 ；





### 2、快速排序实现2

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

void Swap(int *a, int *b) {
    int tmp = *a;
    *a = *b;
    *b = tmp;
}

int Partition(int a[], int l, int r){
    int i, j; 
    int idx = l + rand() % (r - l + 1);        // (1) 
    Swap(&a[l], &a[idx]);                      // (2) 
    i = l, j = r;                              // (3) 
    int x = a[i];                              // (4) 
    while(i < j){
        while(i < j && a[j] > x)               // (5) 
            j--;
        if(i < j) 
             Swap(&a[i], &a[j]), ++i;          // (6) 
                                               // 
                                            
        while(i < j && a[i] < x)               // (7) 
            i++;
        if(i < j)
            Swap(&a[i], &a[j]), --j;           // (8) 
    }
    a[i] = x;
    return i;
}
//递归进行划分
void QuickSort(int a[], int l, int r){
    if(l < r){
        int mid = Partition(a, l, r);
        QuickSort(a, l, mid-1);
        QuickSort(a, mid+1, r);
    }
}

int main() {
    int n;
    while(scanf("%d", &n) != EOF) {
        Input(n, a);
        QuickSort(a, 0, n-1);
        Output(n, a);
    }
    return 0;
} 
```



(1) 随机选择一个基准 ；

(2) 将基准和最左边的值交换 ；

(3) 初始区间 [l, r]；

(4) 这里的 x 是一开始随机出来那个基准值 ；

(5) 从右往左找，第一个满足 a[j] <= x 基准值 的，退出循环；

(6) 如果 a[j] <= x 基准值 ，则 a[j] 必须和 x 交换，才能满足 a[j] 在基准值左边；且此时基准值是 a[i] ，交换完，基准值变成原先的 a[j] ，且 i 需要自增一次； 

(7) 从左往右找，第一个满足 a[i] >= x 基准值 的，退出循环 ；

(8) 如果 a[i] >= x，则 a[i] 必须和 x 交换，才能满足 a[i] 在基准值右边；交换完，基准值又变回为 a[i]，继续迭代； 





### Java版

```java
package Java.DataStructuresAndAlgorithms.Algorithms.sort;

import java.text.SimpleDateFormat;
//import java.util.Arrays;
import java.util.Date;

public class QuickSort {

    public static void main(String[] args) {

        // int[] arr = { -9, 78, 0, 23, -567, 70 };
        // quickSort(arr, 0, arr.length - 1);
        // System.out.println("arr="+Arrays.toString(arr));

        int[] arr = new int[80000];
        for (int i = 0; i < arr.length; i++) {
            arr[i] = (int) (Math.random() * 80000);
        }

        Date date1 = new Date();
        SimpleDateFormat simpleDateFormat = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
        String date1Str = simpleDateFormat.format(date1);
        System.out.println("排序前的时间是=" + date1Str);

        quickSort(arr, 0, arr.length - 1);

        Date date2 = new Date();
        // SimpleDateFormat simpleDateFormat = new SimpleDateFormat("yyyy-MM-dd
        // HH:mm:ss");
        String date2Str = simpleDateFormat.format(date2);
        System.out.println("排序后的时间是=" + date2Str);
    }

    public static void quickSort(int[] arr, int left, int right) {
        int l = left;
        int r = right;
        // pivot是中轴值
        int pivot = arr[(left + right) / 2];
        int temp = 0;
        // while就是让我们的比这个pivot小的放在左边，大的放在右边
        while (l < r) {
            // 在pivot的左边一直去找，找到大于等于的才去退出
            while (arr[l] < pivot) {
                l += 1;
            }
            while (arr[r] > pivot) {
                r -= 1;
            }

            if (l >= r) {
                break;
            }

            temp = arr[l];
            arr[l] = arr[r];
            arr[r] = temp;

            if (arr[l] == pivot) {
                r -= 1;
            }
            if (arr[r] == pivot) {
                l += 1;
            }

        }

        if (l == r) {
            l += 1;
            r -= 1;
        }
        if (left < r) {
            quickSort(arr, left, r);
        }

        if (right > l) {
            quickSort(arr, l, right);

        }

    }
}

```

