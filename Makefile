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
	openapi-generator generate -g go-experimental --api-package $@ -i $< -o $@ --package-name $(@:pkg/%=%)
	rm -f $@/go.{mod,sum}

clean:
	rm -rf pkg/*
