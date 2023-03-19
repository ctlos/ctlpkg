FROM archlinux/archlinux:base-devel
LABEL maintainer="Dancheg97 <dancheg97.ru>"

RUN pacman -Syu --needed --noconfirm go

ARG user=makepkg
RUN useradd --system --create-home $user
RUN echo "$user ALL=(ALL:ALL) NOPASSWD:ALL" > /etc/sudoers.d/$user
USER $user
WORKDIR /home/$user

RUN git clone https://aur.archlinux.org/yay.git
RUN cd yay && makepkg -sri --needed --noconfirm
RUN cd && rm -rf .cache yay

COPY . /home/$user

ENTRYPOINT ["go"]
CMD ["run ."]