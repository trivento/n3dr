#!/bin/bash -ex

if [ -z "${N3DR_APT_GPG_SECRET}" ]; then
  echo "N3DR_APT_GPG_SECRET should not be empty"
  echo "Create one by running:"
  echo "docker run -v /tmp/gpg-output:/root/.gnupg -v ${PWD}/test/gpg/:/tmp/ --rm -it vladgh/gpg --batch --generate-key /tmp/generate"
  echo "docker run --rm -it -v /tmp/gpg-output:/root/.gnupg -v ${PWD}/test/gpg/:/tmp/ vladgh/gpg --output /tmp/my_rsa_key --armor --export-secret-key joe@foo.bar"
  echo "Enter 'abc' as a password, if the prompt appears"
  printf "export N3DR_APT_GPG_SECRET=\$(sudo cat test/gpg/my_rsa_key | tr '\\\n' ' ' | sed -r \"s|-----[A-Z]+ PGP PRIVATE KEY BLOCK-----||g;s| |\\\\\\\\\\\n|g;s|(.*)|-----BEGIN PGP PRIVATE KEY BLOCK-----\\\1-----END PGP PRIVATE KEY BLOCK-----|g\")"
  echo
  echo "sudo rm -r /tmp/gpg-output"
  echo "rm test/gpg/my_rsa_key"
  echo
  printf "Note: Spaces and enters have to be escaped, i.e. '\\\n'->'\\\\\\\n' and ' '->'\ ' if the token is used in travis."
  exit 1
fi

if [ -z "${NEXUS_VERSION}" ]; then
  echo "NEXUS_VERSION empty, setting it to the default value"
  NEXUS_VERSION=3.42.0
fi

if [ -z "${NEXUS_API_VERSION}" ]; then
  echo "NEXUS_API_VERSION empty, setting it to the default value"
  NEXUS_API_VERSION=v1
fi

if [ -z "${N3DR_DELIVERABLE}" ]; then
  echo "N3DR_DELIVERABLE empty, setting it to the default value"
  N3DR_DELIVERABLE=./n3dr
fi

if [ -z "${N3DR_CLEAN_IN_CASE_OF_SUCCESS_OR_FAILURE}" ]; then
  echo "N3DR_CLEAN_IN_CASE_OF_SUCCESS_OR_FAILURE empty, setting it to the default value"
  N3DR_CLEAN_IN_CASE_OF_SUCCESS_OR_FAILURE=true
fi

readonly DOCKER_REGISTRY_HTTP_CONNECTOR_A=8888
readonly DOCKER_REGISTRY_HTTP_CONNECTOR_B=8887
readonly DOCKER_REGISTRY_HTTP_CONNECTOR_C=8886
readonly DOCKER_REGISTRY_HTTP_CONNECTOR_INTERNAL=8083
readonly DOCKER_URL=http://localhost
readonly DOWNLOAD_LOCATION=/tmp/n3dr
readonly DOWNLOAD_LOCATION_PASS=${DOWNLOAD_LOCATION}-pass
readonly DOWNLOAD_LOCATION_RPROXY=${DOWNLOAD_LOCATION}-rproxy
readonly DOWNLOAD_LOCATION_SYNC=${DOWNLOAD_LOCATION}-sync
readonly DOWNLOAD_LOCATION_SYNC_A=${DOWNLOAD_LOCATION_SYNC}-a
readonly DOWNLOAD_LOCATION_SYNC_B=${DOWNLOAD_LOCATION_SYNC}-b
readonly DOWNLOAD_LOCATION_SYNC_C=${DOWNLOAD_LOCATION_SYNC}-c
readonly DOWNLOAD_LOCATION_SYNC_SIZE=23544
readonly HOSTED_REPO_DOCKER=REPO_NAME_HOSTED_DOCKER
readonly HOSTED_REPO_YUM=REPO_NAME_HOSTED_YUM
readonly PORT_NEXUS_A=9999
readonly PORT_NEXUS_B=9998
readonly PORT_NEXUS_C=9997
readonly URL_NEXUS_A=http://localhost:${PORT_NEXUS_A}
readonly URL_NEXUS_B=http://localhost:${PORT_NEXUS_B}
readonly URL_NEXUS_C=http://localhost:${PORT_NEXUS_C}
readonly URL_NEXUS_A_V2="${URL_NEXUS_A/http:\/\//}"
readonly URL_NEXUS_B_V2="${URL_NEXUS_B/http:\/\//}"
readonly URL_NEXUS_C_V2="${URL_NEXUS_C/http:\/\//}"

validate() {
  if [ -z "${N3DR_DELIVERABLE}" ]; then
    echo "No deliverable defined. Assuming that 'go run main.go' should be run."
    N3DR_DELIVERABLE="go run main.go"
  fi
  if [ -z "${NEXUS_VERSION}" ] || [ -z "${NEXUS_API_VERSION}" ]; then
    echo "NEXUS_VERSION and NEXUS_API_VERSION should be specified."
    exit 1
  fi
  if [ -d "${DOWNLOAD_LOCATION}" ]; then
    echo "Ensure that ${DOWNLOAD_LOCATION} does not exist"
    exit 1
  fi
}

build() {
  # shellcheck disable=SC1091
  source ./scripts/build.sh
  cd cmd/n3dr
}

startNexus() {
  # shellcheck disable=SC1091
  # as nexus-docker.sh is retrieved remotely
  source ./start.sh "${NEXUS_VERSION}" "${NEXUS_API_VERSION}" "nexus-${1}" "${2}" "${DOWNLOAD_LOCATION_PASS}" "${3}" "${DOCKER_REGISTRY_HTTP_CONNECTOR_INTERNAL}" &>/dev/null
}

nexus() {
  curl -sL https://gist.githubusercontent.com/030/666c99d8fc86e9f1cc0ad216e0190574/raw/6dd6ce267cba17139d4da0ece908b8c67c14bd21/nexus-docker.sh -o start.sh
  chmod +x start.sh

  startNexus a ${PORT_NEXUS_A} ${DOCKER_REGISTRY_HTTP_CONNECTOR_A} &
  startNexus b ${PORT_NEXUS_B} ${DOCKER_REGISTRY_HTTP_CONNECTOR_B} &
  startNexus c ${PORT_NEXUS_C} ${DOCKER_REGISTRY_HTTP_CONNECTOR_C} &
  wait

  PASSWORD_NEXUS_A=$(cat ${DOWNLOAD_LOCATION_PASS}/nexus-a.txt)
  readonly PASSWORD_NEXUS_A
  PASSWORD_NEXUS_B=$(cat ${DOWNLOAD_LOCATION_PASS}/nexus-b.txt)
  readonly PASSWORD_NEXUS_B
  PASSWORD_NEXUS_C=$(cat ${DOWNLOAD_LOCATION_PASS}/nexus-c.txt)
  readonly PASSWORD_NEXUS_C
}

artifact() {
  mkdir -p "maven-releases/some/group${1}/File_${1}/1.0.0-2"
  echo someContent >"maven-releases/some/group${1}/File_${1}/1.0.0-2/File_${1}-1.0.0-2.jar"
  echo someContentZIP >"maven-releases/some/group${1}/File_${1}/1.0.0-2/File_${1}-1.0.0-2.zip"
  echo -e " <project >\n <modelVersion >4.0.0 </modelVersion >\n <groupId >some.group${1} </groupId >\n <artifactId >File_${1} </artifactId >\n <version >1.0.0-2 </version >\n </project >" >"maven-releases/some/group${1}/File_${1}/1.0.0-2/File_${1}-1.0.0-2.pom"
}

files() {
  for a in $(seq 100); do artifact "${a}"; done
}

upload() {
  echo " #134 archetype-catalog download issue"
  echo "URL: '${URL_NEXUS_A}/repository/maven-releases/archetype-catalog.xml'"
  echo "does not seem to contain a Maven artifact"
  curl -f ${URL_NEXUS_A}/repository/maven-releases/archetype-catalog.xml

  echo "Testing upload..."
  ${N3DR_DELIVERABLE_WITH_BASE_OPTIONS_A} upload \
    -r maven-releases
  echo
}

createHostedAPT() {
  echo "Creating apt repo..."
  curl "${2}/service/rest/beta/repositories/apt/hosted" \
    -s \
    -f \
    -u "admin:${1}" \
    -H "accept: application/json" \
    -H "Content-Type: application/json" \
    -d "{\"name\":\"REPO_NAME_HOSTED_APT\",\"online\":true,\"proxy\":{\"remoteUrl\":\"http://nl.archive.ubuntu.com/ubuntu/\"},\"storage\":{\"blobStoreName\":\"default\",\"strictContentTypeValidation\":true,\"writePolicy\":\"ALLOW_ONCE\"},\"apt\": {\"distribution\": \"bionic\"},\"aptSigning\": {\"keypair\": \"${N3DR_APT_GPG_SECRET}\",\"passphrase\": \"abc\"}}"
}

uploadDeb() {
  if [ "${NEXUS_API_VERSION}" != "beta" ]; then
    createHostedAPT "${PASSWORD_NEXUS_A}" "${URL_NEXUS_A}"

    mkdir REPO_NAME_HOSTED_APT
    cd REPO_NAME_HOSTED_APT
    curl -sL https://github.com/030/a2deb/releases/download/1.0.0/a2deb_1.0.0-0.deb -o a2deb.deb
    curl -sL https://github.com/030/n3dr/releases/download/5.0.1/n3dr_5.0.1-0.deb -o n3dr.deb
    curl -sL https://github.com/030/informado/releases/download/1.4.0/informado_1.4.0-0.deb -o informado.deb
    cd ..

    echo "Testing deb upload..."
    ${N3DR_DELIVERABLE_WITH_BASE_OPTIONS_A} upload \
      -r=REPO_NAME_HOSTED_APT \
      -t=apt
    echo
  else
    echo "Deb upload not supported in beta API"
  fi
}

createHostedNPM() {
  echo "Creating npm repo..."
  curl "${2}/service/rest/v1/repositories/npm/hosted" \
    -s \
    -f \
    -u "admin:${1}" \
    -H "accept: application/json" \
    -H "Content-Type: application/json" \
    -d "{\"name\":\"REPO_NAME_HOSTED_NPM\",\"online\":true,\"storage\":{\"blobStoreName\":\"default\",\"strictContentTypeValidation\":true,\"writePolicy\":\"ALLOW_ONCE\"}}"
}

createHostedYum() {
  echo "Creating hosted yum repository..."
  ${N3DR_DELIVERABLE_WITH_BASE_OPTIONS} configRepository \
    --configRepoName ${HOSTED_REPO_YUM} \
    --configRepoType yum \
    -p "${1}" \
    -n "${2}" \
    --https=false
}

createHostedDocker() {
  echo "Creating hosted docker repository..."
  ${N3DR_DELIVERABLE_WITH_BASE_OPTIONS} configRepository \
    --configRepoName ${HOSTED_REPO_DOCKER} \
    --configRepoType docker \
    -p "${1}" \
    -n "${2}" \
    --https=false \
    --configRepoDockerPort="${DOCKER_REGISTRY_HTTP_CONNECTOR_INTERNAL}"
}

uploadDocker() {
  if [ "${NEXUS_API_VERSION}" != "beta" ]; then
    local docker_registry_uri=repository/testi/utrecht/n3dr
    createHostedDocker "${PASSWORD_NEXUS_A}" "${URL_NEXUS_A_V2}"

    echo "Testing docker upload..."
    docker login ${DOCKER_URL}:${DOCKER_REGISTRY_HTTP_CONNECTOR_A} \
      -p "${PASSWORD_NEXUS_A}" -u admin

    for d in $(seq 5); do
      local docker_registry_tag=localhost:${DOCKER_REGISTRY_HTTP_CONNECTOR_A}/${docker_registry_uri}:6.${d}.0
      docker pull "utrecht/n3dr:6.${d}.0"
      docker tag "utrecht/n3dr:6.${d}.0" "${docker_registry_tag}"
      docker push "${docker_registry_tag}"
    done

    echo
  else
    echo "docker upload not supported in beta API"
  fi
}

uploadNPM() {
  if [ "${NEXUS_API_VERSION}" != "beta" ]; then
    createHostedNPM "${PASSWORD_NEXUS_A}" "${URL_NEXUS_A}"

    mkdir REPO_NAME_HOSTED_NPM
    cd REPO_NAME_HOSTED_NPM
    curl https://registry.npmjs.org/@babel/core/-/core-7.12.10.tgz -o babel-core.tgz
    cd ..

    echo "Testing NPM upload..."
    ${N3DR_DELIVERABLE_WITH_BASE_OPTIONS_A} upload \
      -r=REPO_NAME_HOSTED_NPM \
      -t=npm
    echo
  else
    echo "NPM upload not supported in beta API"
  fi
}

uploadNuget() {
  if [ "${NEXUS_API_VERSION}" != "beta" ]; then
    mkdir nuget-hosted
    cd nuget-hosted
    curl -sL https://chocolatey.org/api/v2/package/n3dr/5.2.6 -o n3dr.5.2.6.nupkg
    cd ..

    echo "Testing nuget upload..."
    ${N3DR_DELIVERABLE_WITH_BASE_OPTIONS_A} upload \
      -r=nuget-hosted \
      -t=nuget
    echo
  else
    echo "Nuget upload not supported in beta API"
  fi
}

uploadYumArtifact() {
  curl -X 'POST' \
    ${URL_NEXUS_A}/service/rest/v1/components?repository=${HOSTED_REPO_YUM} \
    -s \
    -f \
    -u "admin:${PASSWORD_NEXUS_A}" \
    -H 'accept: application/json' \
    -H 'Content-Type: multipart/form-data' \
    -F "yum.asset=@${1};type=application/x-rpm" \
    -F "yum.asset.filename=${1}"
}

uploadYum() {
  if [ "${NEXUS_API_VERSION}" != "beta" ]; then
    createHostedYum "${PASSWORD_NEXUS_A}" "${URL_NEXUS_A_V2}"

    mkdir ${HOSTED_REPO_YUM}
    cd ${HOSTED_REPO_YUM}
    for i in $(seq 5 9); do
      curl -sL "https://yum.puppet.com/puppet-release-el-${i}.noarch.rpm" \
        -o "puppet${i}.rpm"
      uploadYumArtifact "puppet${i}.rpm"
    done

    cd ..
    echo
  else
    echo "yum upload not supported in beta API"
  fi
}

backupHelper() {
  if [ "${NEXUS_VERSION}" == "3.9.0" ]; then
    count_downloads 301 "${1}"
    test_zip 132 "${1}"
  else
    count_downloads 401 "${1}"
    test_zip 172 "${1}"
  fi
  cleanup_downloads
}

backupRegexHelper() {
  if [ "${NEXUS_VERSION}" == "3.9.0" ]; then
    count_downloads 4 "${1}"
    test_zip 4 "${1}"
  else
    count_downloads 5 "${1}"
    test_zip 4 "${1}"
  fi
  cleanup_downloads
}

anonymous() {
  echo "Testing backup by anonymous user..."
  local downloadDir="${DOWNLOAD_LOCATION}/anonymous/"
  ${N3DR_DELIVERABLE} backup \
    -n ${URL_NEXUS_A} \
    -r maven-releases \
    -v "${NEXUS_API_VERSION}" \
    -z \
    --anonymous \
    --directory-prefix="${downloadDir}" \
    --directory-prefix-zip="${downloadDir}"
  backupHelper "${downloadDir}"
}

backup() {
  echo "Testing backup..."
  local downloadDir="${DOWNLOAD_LOCATION}/backup/"
  ${N3DR_DELIVERABLE_WITH_BASE_OPTIONS_A} backup \
    -r maven-releases \
    -z \
    --directory-prefix="${downloadDir}" \
    --directory-prefix-zip="${downloadDir}"
  backupHelper "${downloadDir}"
}

regex() {
  echo -e "Testing regex..."
  local downloadDir="${DOWNLOAD_LOCATION}/backup-regex/"
  ${N3DR_DELIVERABLE_WITH_BASE_OPTIONS_A} backup \
    -r maven-releases \
    -x 'some/group42' \
    -z \
    --directory-prefix="${downloadDir}" \
    --directory-prefix-zip="${downloadDir}"
  backupRegexHelper "${downloadDir}"

  local downloadDir="${DOWNLOAD_LOCATION}/repositories-regex/"
  ${N3DR_DELIVERABLE_WITH_BASE_OPTIONS_A} repositories \
    -b \
    -x 'some/group42' \
    -z \
    --directory-prefix="${downloadDir}" \
    --directory-prefix-zip="${downloadDir}"
  backupRegexHelper "${downloadDir}"
}

repositories() {
  local cmd="${N3DR_DELIVERABLE_WITH_BASE_OPTIONS_A} repositories"

  echo "Testing repositories..."
  $cmd -a | grep maven-releases

  echo "> Counting number of repositories..."
  expected_number=9
  if [ "${NEXUS_API_VERSION}" == "beta" ]; then
    expected_number=5
  fi
  actual_number="$($cmd -c | tail -n1)"
  echo -n "Number of repositories. Expected: ${expected_number}. Actual: ${actual_number}"
  [ "${actual_number}" == "${expected_number}" ]

  echo "> Testing zip functionality..."
  testZipSizeDir=${DOWNLOAD_LOCATION}/testZipSize/
  $cmd -b -z \
    --directory-prefix ${testZipSizeDir} \
    --directory-prefix-zip ${testZipSizeDir}
  if [ "${NEXUS_VERSION}" == "3.9.0" ]; then
    count_downloads 301 ${testZipSizeDir}
    test_zip 132 ${testZipSizeDir}
  else
    count_downloads 402 ${testZipSizeDir}
    test_zip 212 ${testZipSizeDir}
  fi
  cleanup_downloads
}

zipName() {
  echo "Testing zipName..."
  ${N3DR_DELIVERABLE_WITH_BASE_OPTIONS_A} backup -r=maven-releases -z -i=helloZipFile.zip
  ${N3DR_DELIVERABLE_WITH_BASE_OPTIONS_A} repositories -b -z -i=helloZipRepositoriesFile.zip
  find . -name "helloZip*" -type f | wc -l | grep 2
}

clean() {
  for r in a b c; do docker stop nexus-${r} || true; done
  docker stop rproxy-nginx-nexus3 || true
  cleanup_downloads
}

count_downloads() {
  local actual
  actual=$(find "${2}" -type f | wc -l)
  echo "Expected number of artifacts: ${1}"
  echo "Actual number of artifacts: ${actual}"
  echo "${actual}" | grep "${1}"
}

test_zip() {
  local size
  size=$(du "${2}"*n3dr-backup-*zip)
  echo "Actual ZIP size: ${size}"
  echo "Expected ZIP size: ${1}"
  echo "${size}" | grep "^${1}"
}

cleanup_downloads() {
  rm -rf nuget-hosted
  rm -rf REPO_NAME_HOSTED_APT
  rm -rf REPO_NAME_HOSTED_NPM
  rm -rf ${HOSTED_REPO_YUM}
  rm -rf maven-releases
  rm -rf "${DOWNLOAD_LOCATION}"
  rm -rf "${DOWNLOAD_LOCATION_PASS}"
  rm -rf "${DOWNLOAD_LOCATION_RPROXY}"
  rm -rf "${DOWNLOAD_LOCATION_SYNC}"
  rm -rf "${DOWNLOAD_LOCATION_SYNC_A}"
  rm -rf "${DOWNLOAD_LOCATION_SYNC_B}"
  rm -rf "${DOWNLOAD_LOCATION_SYNC_C}"
  rm -f n3dr-backup-*zip
  rm -f helloZip*zip
}

version() {
  echo "Check whether ./n3dr (N3DR_DELIVERABLE: ${N3DR_DELIVERABLE}) --version returns version"
  ${N3DR_DELIVERABLE} --version
  echo
}

test_sync() {
  ${N3DR_DELIVERABLE_WITH_BASE_OPTIONS} repositoriesV2 \
    --backup \
    --directory-prefix "${1}" \
    -p "${2}" \
    -n "${3}" \
    --https=false \
    --dockerHost="${DOCKER_URL}" \
    --dockerPort="${4}" \
    --dockerPortSecure=false
  size=$(du -s --exclude=*p2iwd* "${1}")
  echo "Backup size '${1}': '${size}'"
  echo "${size}" | grep "${DOWNLOAD_LOCATION_SYNC_SIZE}"

  find "${1}/p2iwd/" -type f | wc -l | grep 45
}

sync() {
  if [ "${NEXUS_API_VERSION}" != "beta" ]; then
    echo "Testing the sync between NexusA and followers NexusB and C..."
    createHostedAPT "${PASSWORD_NEXUS_B}" "${URL_NEXUS_B}"
    createHostedAPT "${PASSWORD_NEXUS_C}" "${URL_NEXUS_C}"
    createHostedDocker "${PASSWORD_NEXUS_B}" "${URL_NEXUS_B_V2}"
    createHostedDocker "${PASSWORD_NEXUS_C}" "${URL_NEXUS_C_V2}"
    createHostedNPM "${PASSWORD_NEXUS_B}" "${URL_NEXUS_B}"
    createHostedNPM "${PASSWORD_NEXUS_C}" "${URL_NEXUS_C}"
    createHostedYum "${PASSWORD_NEXUS_B}" "${URL_NEXUS_B_V2}"
    createHostedYum "${PASSWORD_NEXUS_C}" "${URL_NEXUS_C_V2}"

    ${N3DR_DELIVERABLE_WITH_BASE_OPTIONS} sync \
      --otherNexus3Passwords="${PASSWORD_NEXUS_B}","${PASSWORD_NEXUS_C}" \
      --otherNexus3Users=admin,admin \
      --otherNexus3URLs="${URL_NEXUS_B_V2}","${URL_NEXUS_C_V2}" \
      --directory-prefix "${DOWNLOAD_LOCATION_SYNC}" \
      -p "${PASSWORD_NEXUS_A}" \
      -n "${URL_NEXUS_A_V2}" \
      --dockerHost="${DOCKER_URL}" \
      --dockerPort="${DOCKER_REGISTRY_HTTP_CONNECTOR_A}" \
      --dockerPortSecure=false \
      --otherDockerPorts="${DOCKER_REGISTRY_HTTP_CONNECTOR_B}","${DOCKER_REGISTRY_HTTP_CONNECTOR_C}" \
      --otherDockerHosts=http://localhost,http://localhost \
      --otherDockerSecurePorts=false,false

    test_sync "${DOWNLOAD_LOCATION_SYNC_A}" "${PASSWORD_NEXUS_A}" "${URL_NEXUS_A_V2}" "${DOCKER_REGISTRY_HTTP_CONNECTOR_A}"
    test_sync "${DOWNLOAD_LOCATION_SYNC_B}" "${PASSWORD_NEXUS_B}" "${URL_NEXUS_B_V2}" "${DOCKER_REGISTRY_HTTP_CONNECTOR_B}"
    test_sync "${DOWNLOAD_LOCATION_SYNC_C}" "${PASSWORD_NEXUS_C}" "${URL_NEXUS_C_V2}" "${DOCKER_REGISTRY_HTTP_CONNECTOR_C}"
  else
    echo "RepositoriesV2 sync not supported in beta API"
  fi
}

rproxy() {
  if [ "${NEXUS_API_VERSION}" != "beta" ]; then
    local rproxy_conf=../../test/rproxy-nginx-nexus3.conf
    local rproxy_conf_tmp="${rproxy_conf}.tmp"

    echo "Testing rproxy in front of a nexus server..."
    ip_nexus_a=$(docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' nexus-a)
    mkdir -p ${DOWNLOAD_LOCATION_RPROXY}
    sed -e "s|WILL_BE_REPLACED|${ip_nexus_a}|" "${rproxy_conf}" >"${rproxy_conf_tmp}"
    docker run -d --rm --name rproxy-nginx-nexus3 -p 9990:80 -v "${PWD}"/"${rproxy_conf_tmp}":/etc/nginx/nginx.conf nginx:1.21.5-alpine
    sleep 10
    ${N3DR_DELIVERABLE_WITH_BASE_OPTIONS} repositoriesV2 \
      --count \
      --directory-prefix "${DOWNLOAD_LOCATION_RPROXY}" \
      -p "${PASSWORD_NEXUS_A}" \
      -n localhost:9990 \
      --https=false \
      --basePathPrefix=alternativeBasePathNexus3
  else
    echo "Rproxy check skipped in conjunction with beta API"
  fi
}

main() {
  validate
  build
  nexus
  files

  export PASSWORD=${PASSWORD_NEXUS_A}
  readonly N3DR_DELIVERABLE_WITH_BASE_OPTIONS="${N3DR_DELIVERABLE} -u admin --showLogo=false"
  readonly N3DR_DELIVERABLE_WITH_BASE_OPTIONS_A="${N3DR_DELIVERABLE_WITH_BASE_OPTIONS} -p ${PASSWORD} -n ${URL_NEXUS_A} -v ${NEXUS_API_VERSION}"
  upload
  anonymous
  backup
  uploadDeb
  uploadDocker
  uploadNPM
  uploadNuget
  uploadYum
  repositories
  regex
  zipName
  version
  sync
  rproxy
  bats --tap ../../test/tests.bats
  echo "
In order to debug, issue:
N3DR_CLEAN_IN_CASE_OF_SUCCESS_OR_FAILURE=false ./test/integration-tests.sh and
login to ${URL_NEXUS_A}, ${URL_NEXUS_B} or ${URL_NEXUS_C} and login as admin
and respectively ${PASSWORD_NEXUS_A}, ${PASSWORD_NEXUS_B} or
${PASSWORD_NEXUS_C}"
}

if "${N3DR_CLEAN_IN_CASE_OF_SUCCESS_OR_FAILURE}"; then
  trap clean EXIT
fi

main
