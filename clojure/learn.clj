(ns learn
  (:gen-class)
  (:require [clojure.set :as set]))
;(defn fun-name [args] ())
;(def var-name var-value)
;(operator operand1 operand2 operandn)
(defn basic-func [name]
  (def s (str "Hello " name))
  (def d 1.0)
  (println s (type s) (type d))
  (println (dec d) (inc d))
  (println (max 1 0) (min 1 0) (rem 3 2))
  (println (= d d) (not= d d))
  (println (> d d) (>= d d) (< d d) (<= d d))
  (println (or true true) (and true false) (not true))
  (println (bit-and 00111100 00001101) (bit-or 00111100 00001101) (bit-xor 00111100 00001101) (bit-not 00111100))
  )
;(while(expression) (do codeblock))
;(doseq (sequence) statement#1)
;(dotimes (variable value) statement)
;(loop [binding] (condition (statement) (recur (binding))))
(defn cycle-func []
  (println "while")
  (def x (atom 1))                                          ;x type is atom @x type is int
  (while (< @x 5)
    (do
      (println @x)
      (swap! x inc)))
  (println "doseq")
  (doseq [n [0 1 2]]
    (println n))
  (println "dotimes")
  (dotimes [n 5]
    (println n))
  (println "loop")
  (loop [x 10]
    (when (> x 1)
      (println x)
      (recur (- x 2))))
  )
;(if (condition) statement#1 statement #2)
;(if (condition) (statement #1 statement #1.1) (statement #2 statement #2.1))
;(case expression value1 statement #1 value2 statement #2 valueN statement #N statement #Default)
;(cond (expression evaluation1) statement #1 (expression evaluation2) statement #2 (expression evaluationN) statement #N :else statement #Default)
(defn switch-func []
  (if (= 2 2)
    (println "Values are equal")
    (println "Values aren't equal"))
  (if (not= 2 2)
    (do
      (print "Both the values are equal")
      (println "true"))
    (do
      (print "Both the values aren't equal")
      (println "false")))
  (def x 5)
  (case x
    5 (println "x is 5")
    10 (println "x is 10")
    (println "x is neither 5 nor 10"))
  (cond
    (= x 5) (println "x is 5")
    (= x 10) (println "x is 10")
    :else (println "x is not defined"))
  )
(defn func-func [& params]
  (def s (clojure.string/join " " params))
  (fn name [name] (println (str name s)))
  )
(defn string-func []
  (println (str "l" "o" "l"))
  (println (format "int is%d;str is %s;" 0 "."))
  (println (count "0123"))
  (println (subs "0123" 0 2))
  (println (compare "a" "z"))
  (println (clojure.string/lower-case "ABC"))
  (println (clojure.string/upper-case "abc"))
  (println (clojure.string/join ", " [1 2 3]))
  (println (clojure.string/split "Hello0World" #"\d"))
  (println (clojure.string/split-lines "Hello\nWorld"))
  (println (reverse "Hello World"))
  (println (clojure.string/replace "The tutorial is about Kotlin" #"Kotlin" "Clojure"))
  (println (clojure.string/trim " White spaces "))
  (println (clojure.string/triml " White spaces "))
  (println (clojure.string/trimr " White spaces "))
  )
(defn list-func []
  (println (list 1 2 3 4))                                  ;创建列表
  (println (list* 1 [2, 3]))                                ;创建一个新列表,其中包含其他项目,其中最后一个将被视为序列
  (println (first (list 1 2 3)))                            ;此函数返回列表中的第一项
  (println (nth (list 1 2 3) 1))                            ;此函数返回列表中的第n项
  (println (cons 1 (list 1 2 3)))                           ;返回一个新列表,其中元素被添加到列表的开头
  (println (conj (list 1 2 3) 4 5))                         ;返回一个新列表,其中列表在开头,要追加的元素放在末尾
  (println (rest (list 1 2 3)))                             ;返回列表中第一个项目之后的剩余项目
  )
(defn set-func []
  (println (set '(1 1 2 2)))                                ;创建集合
  (println (sorted-set 3 2 1))                              ;返回一个有序的集合
  (println (get (set '(3 2 1)) 2))                          ;返回索引位置处的元素
  (println (contains? (set '(3 2 1)) 2))                    ;找出集合是否包含某个元素
  (println (conj (set '(3 2 1)) 5))                         ;将一个元素附加到集合并返回新的元素集
  (println (disj (set '(3 2 1)) 2))                         ;删除已有项目
  (println (set/union #{1 2} #{3 4}))                       ;联合两个集合
  (println (set/difference #{1 2} #{2 3}))                  ;返回一个没有剩余集合的元素的第一个集合
  (println (set/intersection #{1 2} #{2 3}))                ;返回一组是输入集的交集
  (println (set/subset? #{1 2} #{2 3}))                     ;判断set1是否为set2的一个子集
  (println (set/superset? #{1 2} #{1 2 3}))                 ;判断set1是否为set2的超集
  )
(defn vector-func []
  (println (vector 1 2 3))                                  ;创建载体
  (println (vector-of :int 1 2 3))                          ;创建一个单一基本类型't'的新向量
  (println (nth (vector 1 2, 3) 0))                         ;此函数返回载体中第n个位置的项目
  (println (get (vector 3 2 1) 2))                          ;返回向量中索引位置处的元素
  (println (conj (vector 3 2 1) 5))                         ;向该载体追加一个元素,并返回一组新的载体元素
  (println (pop (vector 3 2 1)))                            ;返回一个没有最后一个项目的新载体
  (println (subvec (vector 1 2 3 4 5 6 7) 2 5))             ;从开始和结束索引返回子载体
  )
(defn seq-func []
  (println (seq [1 2 3]))                                   ;创建序列
  (println (cons 0 (seq [1 2 3])))                          ;返回一个新序列,其中'x'是第一个元素
  (println (conj [1 2 3] 4))                                ;返回一个新序列,其中“x”是添加到序列末尾的元素
  (println (concat (seq [1 2]) (seq [3 4])))                ;这用于将两个序列连接在一起
  (println (distinct (seq [1 1 2 2])))                      ;用于仅确保将不同的元素添加到序列
  (println (reverse (seq [1 2 3])))                         ;反转序列中的元素
  (println (first (seq [1 2 3])) (last (seq [1 2 3])))      ;first last
  (println (rest (seq [1 2 3])))                            ;返回除第一个元素之外的整个序列
  (println (sort (seq [3 2 1])))                            ;排序
  (println (drop 2 (seq [3 2 1])))                          ;从基于需要删除的元素数量的序列中删除元素
  (println (take 2 (seq [3 2 1])))                          ;从序列中获取元素的第一个列表
  (println (take-last 2 (seq [3 2 1])))                     ;从序列中获取元素的最后一个列表
  (println (split-at 2 (seq [1 1 2 2])))                    ;将项目序列拆分为两部分,指定拆分应发生的位置
  )
(defn map-func []
  (def demo0 (hash-map "z" "1" "b" "2" "a" "3"))
  (def demo1 (sorted-map "z" "1" "b" "2" "a" "3"))
  (println demo0 demo1)
  (println (get demo0 "b"))
  (println (contains? demo0 "b") (contains? demo0 "c"))
  (println (find demo0 "b") (find demo0 "c"))
  (println (keys demo0) (vals demo0))
  (println (dissoc demo0 "b"))
  (println (merge-with + (hash-map "z" 1 "b" 2 "a" 3) (hash-map "a" 2 "h" 5 "i" 7)))
  (println (select-keys (hash-map "z" 1 "b" 2 "a" 3) ["z" "a"]))
  (println (set/rename-keys (hash-map "z" 1 "b" 2 "a" 3) {"z" "newz" "b" "newb" "a" "newa"}))
  (println (set/map-invert demo1))
  )
(defn namespace-func []
  (println *ns*)                                            ;这用于查看当前的命名空间
  (println (all-ns))                                        ;所有命名空间
  (println (find-ns 'clojure.string))                       ;查找命名空间
  (println (ns-name 'clojure.string))                       ;获得命名空间的名称
  (println (ns-aliases 'clojure.core))                      ;返回与任何命名空间关联的别名
  (println (count (ns-map 'clojure.core)))                  ;返回命名空间的所有映射的映射
  (alias 'l 'clojure.core)                                  ;alias别名
  (ns-unalias 'clojure.core 'l)                             ;删除alias别名
  )
(defn exception-func []
  (try
    (aget (int-array [1 2 3]) 5)
    (catch Exception e (println (str "caught exception: " (.toString e))))
    (finally (println "This is our final block")))
  (try
    (slurp "xxx.file")
    (catch java.io.IOException e (println (str "caught ioexception: " (.toString e))))
    (catch Exception e (println (str "caught exception: " (.toString e)))))
  )
(defn regular-func []
  (def pat (re-pattern "\\d+"))
  (println (re-find pat "abc123de"))
  (println (clojure.string/replace "abc123de123" pat "789"))
  (println (clojure.string/replace-first "abc123de123" pat "789"))
  )
(defn predicates-func []
  (println (even? 0) (neg? 2) (odd? 3) (pos? 3))
  (println ((every-pred number? even?) 2 4 6))
  (println ((every-pred number? odd?) 2 4 6))
  (println (every? even? '(2 4 6)))
  (println (every? odd? '(2 4 6)))
  (println (some even? '(1 2 3 4)))
  (println (not-any? even? '(2 4 6)))
  )
(defn deconstruction-func []
  (def my-vector [1 2 3 4])
  (let [[a b c d] my-vector] (println a b c d))
  (let [[a b c d e] my-vector] (println a b c d e))
  (let [[a b & the-rest] my-vector] (println a b the-rest))
  (def my-map {"a" 1 "b" 2})
  (let [{a "a" b "b"} my-map] (println a b))
  (let [{a "a" b "b" c "c"} my-map] (println a b c))
  )
(defn file-func []
  (def string1 (slurp "main.clj"))
  (println string1)
  (with-open [rdr (clojure.java.io/reader "main.clj")]
    (reduce conj [] (line-seq rdr)))
  )

(basic-func "cj")
(cycle-func)
(switch-func)
((func-func "a" "b" "c") ":")
(string-func)
(list-func)
(set-func)
(vector-func)
(seq-func)
(map-func)
(namespace-func)
(exception-func)
(regular-func)
(predicates-func)
(deconstruction-func)
;(file-func)
