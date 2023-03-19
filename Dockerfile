FROM golang:latest AS go-build

WORKDIR /src
COPY go.mod /src
COPY go.sum /src
RUN go mod download

COPY . /src
RUN go build -o go-pacman ./main.go

FROM cirrusci/flutter AS flutter-build

WORKDIR /src
COPY . /src

RUN flutter clean
RUN flutter build web

FROM archlinux/archlinux:base-devel
LABEL maintainer="Dancheg97 <dangdancheg@gmail.com>"

RUN pacman -Syu --needed --noconfirm git

ARG user=makepkg
RUN useradd --system --create-home $user
RUN echo "$user ALL=(ALL:ALL) NOPASSWD:ALL" > /etc/sudoers.d/$user
USER $user
WORKDIR /home/$user

RUN git clone https://aur.archlinux.org/yay.git
RUN cd yay && makepkg -sri --needed --noconfirm
RUN cd && rm -rf .cache yay

COPY --from=go-build /src/go-pacman .
COPY --from=flutter-build /src/build/web /web
RUN sed -i 's,href="/,href="/web/,g' /src/build/web/index.html

ENTRYPOINT ["./go-pacman"]
CMD ["--help"]