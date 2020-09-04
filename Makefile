APIS=accounts_and_users \
     stacks \
	 sites \
	 cdn \
	 waf \
	 dns \
	 edge_compute \
	 edge_compute_networking \
	 ssl \
	 monitoring \
	 object_storage

all: $(APIS:%=pkg/%)

pkg/%: oas/%.json
	openapi-generator generate -g go-experimental --api-package $@ -i $< -o $@ \
	               --package-name $(@:pkg/%=%) \
				   --additional-properties enumClassPrefix=true,isGoSubmodule=true
	rm -f $@/go.{mod,sum}

clean:
	for api in ${APIS}; do \
		rm -rf pkg/$${api}; \
	done
