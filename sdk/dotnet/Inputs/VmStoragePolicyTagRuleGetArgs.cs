// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.VSphere.Inputs
{

    public sealed class VmStoragePolicyTagRuleGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to include datastores with the given tags or exclude. Default 
        /// value is true i.e. include datastores with the given tags.
        /// </summary>
        [Input("includeDatastoresWithTags")]
        public Input<bool>? IncludeDatastoresWithTags { get; set; }

        /// <summary>
        /// Name of the tag category.
        /// </summary>
        [Input("tagCategory", required: true)]
        public Input<string> TagCategory { get; set; } = null!;

        [Input("tags", required: true)]
        private InputList<string>? _tags;

        /// <summary>
        /// List of Name of tags to select from the given category.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public VmStoragePolicyTagRuleGetArgs()
        {
        }
    }
}