modules_dir=$PWD/src/modules

module_str=''
import_str='"github.com/rainkid/dogo"\n'

buidmodules () {
	for file in `find $modules_dir -type f`
	    do
			module=`echo $(dirname "$file")|xargs basename`
		    controller=`echo $(basename $file)|cut -d'.' -f1`
			first=`echo $controller| cut -c1 | tr '[:lower:]' '[:upper:]'`
			last=`echo $controller| cut -c2-${#controller} | tr '[:upper:]' '[:lower:]'`
			controller=$first$last
			if [ "Base" != "$controller" ]; then
				module_str=$module_str'router.AddSampleRoute("'$module'",&'$module'.'$controller'{})\n'
			fi
	    done
	
	for dir in `find $modules_dir -type d` 
	    do 
	    	if [ "$modules_dir" != "$dir" ]; then
				count=`ls $dir|wc -l`
				if [ 0 -ne $count ]; then
					module=`echo $(basename "$dir")`
					import_str=$import_str$module' "modules/'$module'"'
				fi
			fi
	    done
}

buidmodules
echo 'package main\n
import ('$import_str')\n
func AddSampleRoute(router *dogo.Router) {\n
	'$module_str'\n
}\n' > routers.go

go fmt routers.go
