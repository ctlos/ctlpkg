FROM golang:latest AS go-build
WORKDIR /
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o go-pacman ./main.go

FROM instrumentisto/flutter AS flutter-build
WORKDIR /
COPY . .
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

COPY --from=go-build /go-pacman .
COPY --from=flutter-build /build/web /web

ENTRYPOINT ["./go-pacman"]
CMD ["--help"]