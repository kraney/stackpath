APIS=accounts-and-users \
     stacks \
	 sites \
	 cdn \
	 waf \
	 dns \
	 edge-compute \
	 edge-compute-networking \
	 ssl \
	 monitoring \
	 object-storage

all: $(APIS:%=pkg/%)

pkg/%: oas/%.json
	openapi-generator generate -g go-experimental --api-package $@ -i $< -o $@ --package-name $(@:pkg/%=%)
	rm -f pkg/$@/go.{mod,sum}

clean:
	rm -rf pkg/*
