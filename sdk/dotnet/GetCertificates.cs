// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Octopusdeploy
{
    public static class GetCertificates
    {
        /// <summary>
        /// Provides information about existing certificates.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Octopusdeploy = Pulumi.Octopusdeploy;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Octopusdeploy.GetCertificates.Invoke(new()
        ///     {
        ///         Archived = "false",
        ///         Ids = new[]
        ///         {
        ///             "Certificates-123",
        ///             "Certificates-321",
        ///         },
        ///         PartialName = "Defau",
        ///         Skip = 5,
        ///         Take = 100,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCertificatesResult> InvokeAsync(GetCertificatesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCertificatesResult>("octopusdeploy:index/getCertificates:getCertificates", args ?? new GetCertificatesArgs(), options.WithDefaults());

        /// <summary>
        /// Provides information about existing certificates.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Octopusdeploy = Pulumi.Octopusdeploy;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Octopusdeploy.GetCertificates.Invoke(new()
        ///     {
        ///         Archived = "false",
        ///         Ids = new[]
        ///         {
        ///             "Certificates-123",
        ///             "Certificates-321",
        ///         },
        ///         PartialName = "Defau",
        ///         Skip = 5,
        ///         Take = 100,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetCertificatesResult> Invoke(GetCertificatesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCertificatesResult>("octopusdeploy:index/getCertificates:getCertificates", args ?? new GetCertificatesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCertificatesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A filter to search for resources that have been archived.
        /// </summary>
        [Input("archived")]
        public string? Archived { get; set; }

        [Input("certificates")]
        private List<Inputs.GetCertificatesCertificateArgs>? _certificates;

        /// <summary>
        /// A list of certificates that match the filter(s).
        /// </summary>
        public List<Inputs.GetCertificatesCertificateArgs> Certificates
        {
            get => _certificates ?? (_certificates = new List<Inputs.GetCertificatesCertificateArgs>());
            set => _certificates = value;
        }

        /// <summary>
        /// A filter to define the first result.
        /// </summary>
        [Input("firstResult")]
        public string? FirstResult { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A filter to search by a list of IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A filter used to order the search results.
        /// </summary>
        [Input("orderBy")]
        public string? OrderBy { get; set; }

        /// <summary>
        /// A filter to search by the partial match of a name.
        /// </summary>
        [Input("partialName")]
        public string? PartialName { get; set; }

        /// <summary>
        /// A filter of terms used the search operation.
        /// </summary>
        [Input("search")]
        public string? Search { get; set; }

        /// <summary>
        /// A filter to specify the number of items to skip in the response.
        /// </summary>
        [Input("skip")]
        public int? Skip { get; set; }

        /// <summary>
        /// A filter to specify the number of items to take (or return) in the response.
        /// </summary>
        [Input("take")]
        public int? Take { get; set; }

        /// <summary>
        /// A filter to search by a tenant ID.
        /// </summary>
        [Input("tenant")]
        public string? Tenant { get; set; }

        public GetCertificatesArgs()
        {
        }
        public static new GetCertificatesArgs Empty => new GetCertificatesArgs();
    }

    public sealed class GetCertificatesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A filter to search for resources that have been archived.
        /// </summary>
        [Input("archived")]
        public Input<string>? Archived { get; set; }

        [Input("certificates")]
        private InputList<Inputs.GetCertificatesCertificateInputArgs>? _certificates;

        /// <summary>
        /// A list of certificates that match the filter(s).
        /// </summary>
        public InputList<Inputs.GetCertificatesCertificateInputArgs> Certificates
        {
            get => _certificates ?? (_certificates = new InputList<Inputs.GetCertificatesCertificateInputArgs>());
            set => _certificates = value;
        }

        /// <summary>
        /// A filter to define the first result.
        /// </summary>
        [Input("firstResult")]
        public Input<string>? FirstResult { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A filter to search by a list of IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A filter used to order the search results.
        /// </summary>
        [Input("orderBy")]
        public Input<string>? OrderBy { get; set; }

        /// <summary>
        /// A filter to search by the partial match of a name.
        /// </summary>
        [Input("partialName")]
        public Input<string>? PartialName { get; set; }

        /// <summary>
        /// A filter of terms used the search operation.
        /// </summary>
        [Input("search")]
        public Input<string>? Search { get; set; }

        /// <summary>
        /// A filter to specify the number of items to skip in the response.
        /// </summary>
        [Input("skip")]
        public Input<int>? Skip { get; set; }

        /// <summary>
        /// A filter to specify the number of items to take (or return) in the response.
        /// </summary>
        [Input("take")]
        public Input<int>? Take { get; set; }

        /// <summary>
        /// A filter to search by a tenant ID.
        /// </summary>
        [Input("tenant")]
        public Input<string>? Tenant { get; set; }

        public GetCertificatesInvokeArgs()
        {
        }
        public static new GetCertificatesInvokeArgs Empty => new GetCertificatesInvokeArgs();
    }


    [OutputType]
    public sealed class GetCertificatesResult
    {
        /// <summary>
        /// A filter to search for resources that have been archived.
        /// </summary>
        public readonly string? Archived;
        /// <summary>
        /// A list of certificates that match the filter(s).
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCertificatesCertificateResult> Certificates;
        /// <summary>
        /// A filter to define the first result.
        /// </summary>
        public readonly string? FirstResult;
        /// <summary>
        /// An auto-generated identifier that includes the timestamp when this data source was last modified.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A filter to search by a list of IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// A filter used to order the search results.
        /// </summary>
        public readonly string? OrderBy;
        /// <summary>
        /// A filter to search by the partial match of a name.
        /// </summary>
        public readonly string? PartialName;
        /// <summary>
        /// A filter of terms used the search operation.
        /// </summary>
        public readonly string? Search;
        /// <summary>
        /// A filter to specify the number of items to skip in the response.
        /// </summary>
        public readonly int? Skip;
        /// <summary>
        /// A filter to specify the number of items to take (or return) in the response.
        /// </summary>
        public readonly int? Take;
        /// <summary>
        /// A filter to search by a tenant ID.
        /// </summary>
        public readonly string? Tenant;

        [OutputConstructor]
        private GetCertificatesResult(
            string? archived,

            ImmutableArray<Outputs.GetCertificatesCertificateResult> certificates,

            string? firstResult,

            string id,

            ImmutableArray<string> ids,

            string? orderBy,

            string? partialName,

            string? search,

            int? skip,

            int? take,

            string? tenant)
        {
            Archived = archived;
            Certificates = certificates;
            FirstResult = firstResult;
            Id = id;
            Ids = ids;
            OrderBy = orderBy;
            PartialName = partialName;
            Search = search;
            Skip = skip;
            Take = take;
            Tenant = tenant;
        }
    }
}