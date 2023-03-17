FROM golang:latest AS build
WORKDIR /
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o go-pacman ./main.go

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

COPY --from=build /go-pacman .

ENTRYPOINT ["./go-pacman"]
CMD ["--help"]