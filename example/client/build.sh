
# Setup
sudo apt-get update
sudo apt-get install apt-transport-https
sudo sh -c 'curl https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add -'
sudo sh -c 'curl https://storage.googleapis.com/download.dartlang.org/linux/debian/dart_stable.list > /etc/apt/sources.list.d/dart_stable.list'
sudo apt-get update

# Install
sudo apt-get install dart

cd $GOPATH/src/github.com/platlk/satms/example/client
pub get
pub build

echo 'Now, you can launch the client with your favorite browser'
echo '<your_browser> $GOPATH/src/github.com/platlk/satms/example/client/build/main.html'