gen:
  flatc --go --gen-all \
    --go-namespace protocol \
    -o . \
    protocol/fbs/protocol.fbs
