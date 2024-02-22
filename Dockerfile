# Build Out Provider
FROM golang:1.21-alpine AS compiler
WORKDIR /terraform-provider-unicode
COPY  . .
#RUN rm $GOPATH/go.mod
RUN go build .
RUN chmod +x terraform-provider-unicode
RUN chown  1000:1000 terraform-provider-unicode
RUN chgrp 1000 terraform-provider-unicode

# Temporary Container TO Build User
FROM debian:stretch-slim as auth
RUN groupadd -g 1000 app && useradd -u 1000 -g 1000 -m app

# Copy Into Terraform Container 
FROM hashicorp/terraform:light as default 
# Make User 1000
COPY --from=auth /etc/passwd /etc/passwd
COPY --from=auth /etc/group /etc/group
COPY --from=auth /home/app /home/app
RUN chown 1000:1000 /home/app
# Make Home Directory
COPY --from=compiler /terraform-provider-unicode/terraform-provider-unicode /home/app/terraform-provider-unicode
COPY --from=compiler /terraform-provider-unicode/.terraformrc /home/app/.terraformrc
RUN chown 1000:1000 /home/app/.terraformrc
RUN chown 1000:1000 /home/app/terraform-provider-unicode
RUN chmod +x /home/app/terraform-provider-unicode
# Now Switch to User app
USER app 
WORKDIR /home/app
ENTRYPOINT ["terraform", "apply", "-auto-approve"]
CMD ["--var", "user=default"]
# does this user follow you arround?

# American App -> Take in variables and output
FROM default as american
USER app
WORKDIR /home/app
RUN mkdir america-themed-app
COPY --from=compiler /terraform-provider-unicode/examples/america-themed_app/ /home/app/america-themed-app
WORKDIR /home/app/america-themed-app  
#&& terraform init
ENTRYPOINT [ "terraform", "apply","-auto-approve"]
CMD ["--var", "user=default" ]

FROM default as amphibian
USER app
WORKDIR /home/app
RUN mkdir amphibian-themed-app
COPY --from=compiler /terraform-provider-unicode/examples/amphibian_themed_app/ /home/app/amphibian-themed-app
WORKDIR /home/app/amphibian-themed-app
#&& terraform init
ENTRYPOINT [ "terraform","apply","-auto-approve"]
CMD ["--var", "user=default" ]


FROM default as halloween 
USER app
WORKDIR /home/app
RUN mkdir halloween-themed-app
COPY --from=compiler /terraform-provider-unicode/examples/halloween_themed_app/ /home/app/halloween-themed-app    
WORKDIR /home/app/halloween-themed-app

ENTRYPOINT [ "terraform", "apply","-auto-approve"]
CMD ["--var", "user=default" ]
