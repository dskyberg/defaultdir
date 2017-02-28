# defaultdir
Go utility to determine a default location for a file system folder.  Defaultdir
is a fluent style utility, so it's easy to describe the preference order you
want to use.

## Order of Evaluation
The goal of difaultdir is to make it easy to describe an order of evaluation. 
Thus, once a target has been found, no other processing is done.  So, be sure
to place the methods in the order you want them evaluated!

## Errors
Once defaultdir encounters an error, no more work is done in the fluent chain.
Thus, the error returned from `dir() (string, error)` will always be the error
that broke the chain.

## Examples

```go
func getDefaultDir() string {
   envKey := "MY_DIR"
   base := "base_dir"

	dir, err := defaultdir.New().
		Env(envKey). // Use value of MY_DIR as the directory
		Base(base).  // Set "base_dir" as the sub folder to look for
		Env(envKey). // Now look for "<value of MY_DIR>/base_dir"
		Cwd().       // Use $CWD/base_dir
		Bin().       // Use <app exe dir>/base_dir
      ClearBase(). // Unset "base_dir"
      Cwd().       // If nothing else if found, use $CWD
		Dir()
   
   // There should be no error, and dir should be set
	if err == nil && dir != nil {
		return *dir
	}
   
}
```
