FROM cirrusci/flutter 

LABEL maintainer="Dancheg97 <dancheg97.ru>"

WORKDIR /src
COPY . /src

RUN flutter pub get
RUN flutter build web

ENTRYPOINT [ "python3" ]
CMD [ "-m http.server -d /src/build/web 80" ]
