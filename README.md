<a><img src="https://github.com/user-attachments/assets/1f233586-a2c3-4838-beeb-3fcbd793ba82" height="150" width="150"></a>

![Tests](https://github.com/Greeshmanth1909/Shadowfax/actions/workflows/ci.yml/badge.svg)
# Shadowfax
Shadowfax is a UCI chess engine written entirely from scratch in Golang. It will be using classic evlauation and move generation methods. I also plan on upgrading it to an NNUE in version 2. If that sounds interesting, consider giving this repo a star. Thanks!

# Note
The project is still in the early stages of development and not completely functional.

# testing
- All unit tests must make sure `util.InitAll()` is called before starting. the results will be inconsistent if this is not followed
