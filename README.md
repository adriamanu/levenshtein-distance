# levenshtein-distance
## Levenshtein distance implementation in golang

https://en.wikipedia.org/wiki/Levenshtein_distance

Distance from comit to commit is 1 (one addition)
```
    c o m m i t
 [0 1 2 3 4 5 6]
c[1 0 1 2 3 4 5]
o[2 1 0 1 2 3 4]
m[3 2 1 0 1 2 3]
i[4 3 2 1 1 1 2]
t[5 4 3 2 2 2 1]
```

Distance from push to pull is 2 (2 modifications)
```
     p u l l
  [0 1 2 3 4]
p [1 0 1 2 3]
u [2 1 0 1 2]
s [3 2 1 1 2]
h [4 3 2 2 2]
```

Distance from maison to maison is 0 because words are identical
```
    m a i s o n
 [0 1 2 3 4 5 6]
m[1 0 1 2 3 4 5]
a[2 1 0 1 2 3 4]
i[3 2 1 0 1 2 3]
s[4 3 2 1 0 1 2]
o[5 4 3 2 1 0 1]
n[6 5 4 3 2 1 0]
```

Distance from zèbre to anathème is 7 (replace 3 and add 4 new letters)
```
    a n a t h è m e
 [0 1 2 3 4 5 6 7 8]
z[1 1 2 3 4 5 6 7 8]
è[2 2 2 3 4 5 5 6 7]
b[3 3 3 3 4 5 6 6 7]
r[4 4 4 4 4 5 6 7 7]
e[5 5 5 5 5 5 5 6 7]
```