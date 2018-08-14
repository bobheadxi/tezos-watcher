FROM tezos/tezos:alphanet

COPY ./sandbox.sh /sandbox.sh

RUN sudo apk update && sudo apk add bash

ENTRYPOINT [ "bash", "/sandbox.sh", "1", "--connections", "1" ]
