# 使 用 轻 量 的 alpine 镜 像 作 为 基 础 镜 像
FROM centos:7.9.2009
# 设 置 工 作 目 录
WORKDIR /app
# 复 制 二 进 制 文 件
COPY alert-plugin .
# 暴 露 端 口
EXPOSE 8080
# 运 行 应 用
CMD ["./alert-plugin"]