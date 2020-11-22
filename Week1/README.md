# 学习笔记

## 分析PriorityQueue

- 源码地址
  - http://developer.classpath.org/doc/java/util/PriorityQueue-source.html

- 实现原理
  - 使用了一个binomial heap.
    - binomial heap相关信息，https://zh.wikipedia.org/wiki/%E4%BA%8C%E9%A1%B9%E5%A0%86
    - 主要优势
      - 运行时间，以下对于二项堆操作的运行时间都为O(log n)：
        - 在二项堆中插入新结点
        - 查找最小关键字所在结点
        - 从二项堆中删除最小关键字所在结点
        - 减小给定结点关键字的值
        - 删除给定结点
        - 合并两个二项堆

- 遗憾，没时间细看二项堆，有点小复杂