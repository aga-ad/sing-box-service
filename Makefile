build_service:
	go build -v -trimpath -ldflags "-X 'github.com/sagernet/sing-box/constant.Version=1.12.12' -s -w -buildid=" -tags "with_gvisor,with_quic,with_dhcp,with_wireguard,with_utls,with_acme,with_clash_api,with_tailscale" ./cmd/sing-box-service
