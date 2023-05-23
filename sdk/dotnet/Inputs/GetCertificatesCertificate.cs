// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace PiersKarsenbarg.Octopusdeploy.Inputs
{

    public sealed class GetCertificatesCertificateArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A filter to search for resources that have been archived.
        /// </summary>
        [Input("archived", required: true)]
        public string Archived { get; set; } = null!;

        [Input("certificateData", required: true)]
        private string? _certificateData;

        /// <summary>
        /// The encoded data of the certificate.
        /// </summary>
        public string? CertificateData
        {
            get => _certificateData;
            set => _certificateData = value;
        }

        /// <summary>
        /// Specifies the archive file format used for storing cryptography objects in the certificate. Valid formats are `Der`, `Pem`, `Pkcs12`, or `Unknown`.
        /// </summary>
        [Input("certificateDataFormat", required: true)]
        public string CertificateDataFormat { get; set; } = null!;

        [Input("environments", required: true)]
        private List<string>? _environments;

        /// <summary>
        /// A list of environment IDs associated with this resource.
        /// </summary>
        public List<string> Environments
        {
            get => _environments ?? (_environments = new List<string>());
            set => _environments = value;
        }

        /// <summary>
        /// Indicates if the certificate has a private key.
        /// </summary>
        [Input("hasPrivateKey", required: true)]
        public bool HasPrivateKey { get; set; }

        /// <summary>
        /// The unique ID for this resource.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        /// <summary>
        /// Indicates if the certificate has expired.
        /// </summary>
        [Input("isExpired", required: true)]
        public bool IsExpired { get; set; }

        [Input("issuerCommonName", required: true)]
        public string IssuerCommonName { get; set; } = null!;

        [Input("issuerDistinguishedName", required: true)]
        public string IssuerDistinguishedName { get; set; } = null!;

        [Input("issuerOrganization", required: true)]
        public string IssuerOrganization { get; set; } = null!;

        /// <summary>
        /// The name of this resource.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("notAfter", required: true)]
        public string NotAfter { get; set; } = null!;

        [Input("notBefore", required: true)]
        public string NotBefore { get; set; } = null!;

        [Input("notes", required: true)]
        public string Notes { get; set; } = null!;

        [Input("password", required: true)]
        private string? _password;

        /// <summary>
        /// The password associated with this resource.
        /// </summary>
        public string? Password
        {
            get => _password;
            set => _password = value;
        }

        [Input("replacedBy", required: true)]
        public string ReplacedBy { get; set; } = null!;

        [Input("selfSigned", required: true)]
        public bool SelfSigned { get; set; }

        [Input("serialNumber", required: true)]
        public string SerialNumber { get; set; } = null!;

        [Input("signatureAlgorithmName", required: true)]
        public string SignatureAlgorithmName { get; set; } = null!;

        [Input("subjectAlternativeNames", required: true)]
        private List<string>? _subjectAlternativeNames;
        public List<string> SubjectAlternativeNames
        {
            get => _subjectAlternativeNames ?? (_subjectAlternativeNames = new List<string>());
            set => _subjectAlternativeNames = value;
        }

        [Input("subjectCommonName", required: true)]
        public string SubjectCommonName { get; set; } = null!;

        [Input("subjectDistinguishedName", required: true)]
        public string SubjectDistinguishedName { get; set; } = null!;

        [Input("subjectOrganization", required: true)]
        public string SubjectOrganization { get; set; } = null!;

        [Input("tenantTags", required: true)]
        private List<string>? _tenantTags;

        /// <summary>
        /// A list of tenant tags associated with this resource.
        /// </summary>
        public List<string> TenantTags
        {
            get => _tenantTags ?? (_tenantTags = new List<string>());
            set => _tenantTags = value;
        }

        /// <summary>
        /// The tenanted deployment mode of the resource. Valid account types are `Untenanted`, `TenantedOrUntenanted`, or `Tenanted`.
        /// </summary>
        [Input("tenantedDeploymentParticipation", required: true)]
        public string TenantedDeploymentParticipation { get; set; } = null!;

        [Input("tenants", required: true)]
        private List<string>? _tenants;

        /// <summary>
        /// A list of tenant IDs associated with this resource.
        /// </summary>
        public List<string> Tenants
        {
            get => _tenants ?? (_tenants = new List<string>());
            set => _tenants = value;
        }

        [Input("thumbprint", required: true)]
        public string Thumbprint { get; set; } = null!;

        [Input("version", required: true)]
        public int Version { get; set; }

        public GetCertificatesCertificateArgs()
        {
        }
        public static new GetCertificatesCertificateArgs Empty => new GetCertificatesCertificateArgs();
    }
}
