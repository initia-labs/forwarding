cd proto
buf generate --template buf.gen.gogo.yaml
buf generate --template buf.gen.pulsar.yaml
cd ..

cp -r github.com/initia-labs/forwarding/v2/* ./
cp -r api/noble/forwarding/* api/

rm -rf github.com
rm -rf api/noble
rm -rf noble
