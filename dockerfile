FROM archlinux/archlinux:base-devel
LABEL maintainer="test.cab <git@test.cab>"
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
RUN go install .
CMD regen
