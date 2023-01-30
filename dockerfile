FROM archlinux/archlinux:base-devel

LABEL maintainer="Dancheg97 <dangdancheg@gmail.com>"

RUN pacman -Syu --needed --noconfirm git go

ARG user=makepkg
RUN useradd --system --create-home $user \
    && echo "$user ALL=(ALL:ALL) NOPASSWD:ALL" > /etc/sudoers.d/$user
USER $user
WORKDIR /home/$user

COPY . .

RUN git clone https://aur.archlinux.org/yay.git \
    && cd yay \
    && makepkg -sri --needed --noconfirm \
    && cd \
    && rm -rf .cache yay

RUN go build -buildvcs=false .
ENTRYPOINT ["./go-pacman"]
CMD ["--help"]