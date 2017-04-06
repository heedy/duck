# Quack

*A 'quick' version of duck*

Using reflect in golang is too slow for some high-performance applications.
Quack is an optimized version of duck with a limited set of conversion types
which runs about 10 times faster.

Quack cannot handle struct types, or pointers. It is built for processing
data that comes from marshalled json or the like. It can process:
- All int variants, including uints
- float32 and float64
- string
- []interface{}, []float64, []int64, []string
- map[string]interface{}, map[string]float64, map[string]int64, map[string]string

The API is the same as duck.


```
benchmark                     old ns/op     new ns/op     delta
BenchmarkLtInt-4              100           14.1          -85.90%
BenchmarkLtFloat-4            100           13.2          -86.80%
BenchmarkLtString-4           289           65.9          -77.20%
BenchmarkEqualInt-4           193           10.4          -94.61%
BenchmarkEqualIntFloat-4      146           17.2          -88.22%
BenchmarkEqualString-4        455           57.2          -87.43%
BenchmarkEqualStringInt-4     251           43.3          -82.75%
BenchmarkAddInt-4             131           44.1          -66.34%
BenchmarkAddFloat-4           131           43.1          -67.10%
BenchmarkAddString-4          578           214           -62.98%
BenchmarkMultiplyFloat-4      100           13.1          -86.90%
BenchmarkModFloat-4           101           12.9          -87.23%
BenchmarkGetMap-4             268           14.8          -94.48%
BenchmarkGetArray-4           161           6.67          -95.86%
BenchmarkArrayLength-4        81.4          5.08          -93.76%
BenchmarkMapLength-4          20.7          4.27          -79.37%
BenchmarkIntToInt-4           48.6          6.32          -87.00%
BenchmarkFloatToInt-4         49.8          6.23          -87.49%
BenchmarkStringToInt-4        142           41.5          -70.77%
BenchmarkIntToFloat-4         48.5          6.22          -87.18%
BenchmarkFloatToFloat-4       48.6          5.44          -88.81%
BenchmarkStringToFloat-4      143           40.5          -71.68%
BenchmarkIntToString-4        102           59.6          -41.57%
BenchmarkFloatToString-4      219           176           -19.63%
BenchmarkStringToString-4     67.6          5.44          -91.95%
```


## Why couldn't duck be modified to include this?

It turns out that even including certain functions (like `strconv.ParseFloat`)
already causes type conversion benchmarks to be 5 times slower - **even if the function is not called**.

That is, parsing strings using strconv made everything else slow too! This was worked around by completely eliminating reflect, and creating a fork of ParseFloat that did not have the issue.