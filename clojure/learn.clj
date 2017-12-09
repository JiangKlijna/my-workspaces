(ns learn
    (:gen-class))
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
    (def x (atom 1));x type is atom @x type is int
    (while ( < @x 5 )
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
    (if ( = 2 2)
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
        (= x 10)(println "x is 10")
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
;(file-func)
